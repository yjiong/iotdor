package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"database/sql"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"github.com/yjiong/iotdor/ent"
	"github.com/yjiong/iotdor/internal/datasrc"
)

// Manage logic handle
type Manage struct {
	ctx context.Context
	datasrc.DSrcer
	mia datasrc.SyncMSG
	*sql.DB
	entC     *ent.Client
	redisC   *redis.Client
	iotdName string
	ids      []cron.EntryID
	*cron.Cron
	storageInterval int64
	spool           sync.Pool
}

// MsgHandle ....
func (m *Manage) MsgHandle() {
	log.Infoln("start logic message handle")
	m.runCron()
	for {
		select {
		case <-m.ctx.Done():
			m.DB.Close()
			m.DSrcer.Close()
			m.entC.Close()
			m.redisC.Close()
			return
		case msg := <-m.DataDownChan():
			gwID := msg.Get("sender").MustString()
			gwCtag := msg.Get("ctag").MustString()
			gwSer := msg.Get("api").MustString()
			if cmd, err := msg.Get("cmd").String(); err == nil {
				log.Debugf("receive mqtt data cmd=%s", cmd)
				tstamp, _ := msg.Get("tstamp").Int64()
				tstr := time.Unix(tstamp, 0).Local().Format("2006-01-02 15:04:05")
				switch cmd {
				case AutoUpdata:
					md := getRawMap(msg.Get("data").MustMap(), tstr)
					if devID, ok := md["devid"].(string); ok {
						md["status"] = "normal"
						md["timestamp"] = tstr
						if e, ok := md["error"]; ok {
							md["status"] = e
							delete(md, "error")
						}
						if err = m.redisC.HSet(m.ctx, m.mkRedisKeyPrefix(gwID, devID, DeviceValue), md).Err(); err != nil {
							log.Error(err)
						}
					}
				case GwStat:
					stat := msg.Get("data").MustInt()
					go m.gatewayInfoHandle(gwID, gwCtag, gwSer, stat)
				default:
					if mid, _ := msg.Get("seq").String(); len(mid) > 0 {
						m.mia.HandleReceive(mid, msg.Get("data").Interface())
					}
				}
			}
		}
	}
}

// NewManage ....
func NewManage(ctx context.Context,
	db *sql.DB,
	dsrc datasrc.DSrcer,
	entc *ent.Client,
	redisc *redis.Client,
	iname string,
	interval int64) *Manage {
	m := &Manage{
		ctx:             ctx,
		DB:              db,
		DSrcer:          dsrc,
		entC:            entc,
		redisC:          redisc,
		iotdName:        iname,
		ids:             []cron.EntryID{},
		Cron:            cron.New(cron.WithSeconds()),
		storageInterval: interval,
		spool: sync.Pool{
			New: func() any { return make(map[string]string) },
		},
	}
	m.mInit()
	m.mia = datasrc.NewMsgInteractive()
	return m
}

func (m *Manage) mInit() {
	addGroupIfNotExist(m.ctx, m.entC, m.iotdName)
}

func (m *Manage) mkRedisKeyPrefix(s ...string) (kp string) {
	kp = m.iotdName
	for _, c := range s {
		kp = fmt.Sprintf("%s:%s", kp, c)
	}
	return strings.ToUpper(kp)
}

func getRawMap(md map[string]interface{}, tstamp string) map[string]interface{} {
	for k, v := range md {
		if iv, ok := v.(json.Number); ok {
			md[k] = string(iv)
		}
	}
	md["status"] = "normal"
	md["timestamp"] = tstamp
	if e, ok := md["error"]; ok {
		md["status"] = e
		delete(md, "error")
	}
	return md
}

func (m *Manage) runCron() {
	cStr := fmt.Sprintf("10 */%d * * * *", m.storageInterval)
	EID1, _ := m.Cron.AddFunc(
		cStr,
		m.storageDeviceValue)
	m.ids = []cron.EntryID{EID1}
	m.Cron.Start()
}

func (m *Manage) cronStop() {
	for _, eid := range m.ids {
		m.Remove(eid)
	}
	m.ids = m.ids[0:0]
	log.Infoln("cron schedule removed")
}

func (m *Manage) storageDeviceValue() {
	keys, err := m.redisC.Keys(m.ctx, fmt.Sprintf("*:%s", DeviceValue)).Result()
	if err != nil {
		log.Error(errors.Wrap(err, "get DEVICE_VALUE faild"))
		return
	}
	for _, k := range keys {
		vs := m.spool.Get().(map[string]string)
		vs, _ = m.redisC.HGetAll(m.ctx, k).Result()
		go InsertMap(m.DB, DEVICETABLE, vs)
		m.spool.Put(vs)
	}
}
