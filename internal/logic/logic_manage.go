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
			if cmd, err := msg.Get("cmd").String(); err == nil {
				log.Debugf("receive mqtt data cmd=%s", cmd)
				gwID := msg.Get("sender").MustString()
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
						if err = m.redisC.HSet(m.ctx, m.mkKeyPrefix(gwID, devID, DeviceValue), md).Err(); err != nil {
							log.Error(err)
						}
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
	return m
}

func (m *Manage) mInit() {
	addGroupIfNotExist(m.ctx, m.entC, m.iotdName)
}

func (m *Manage) mkKeyPrefix(s ...string) (kp string) {
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
		m.storateDeviceValue)
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

func (m *Manage) storateDeviceValue() {
	keys, err := m.redisC.Keys(m.ctx, fmt.Sprintf("*:%s", DeviceValue)).Result()
	if err != nil {
		log.Error(errors.Wrap(err, "get DEVICE_VALUE faild"))
	}
	for _, k := range keys {
		vs := m.spool.Get().(map[string]string)
		vs, _ = m.redisC.HGetAll(m.ctx, k).Result()
		go InsertMap(m.DB, "ammeters", vs)
		m.spool.Put(vs)
	}
}

// AllDevices ....
func (m *Manage) AllDevices() (ids []string) {
	keys, _ := m.redisC.Keys(m.ctx, fmt.Sprintf("*:%s", DeviceValue)).Result()
	for _, key := range keys {
		kids := strings.Split(key, ":")
		if len(kids) > 3 {
			ids = append(ids, kids[2])
		}
	}
	return
}

// DeviceRealTimeValue ....
func (m *Manage) DeviceRealTimeValue(devid string) map[string]string {
	keys, _ := m.redisC.Keys(m.ctx, fmt.Sprintf("*:%s:*", devid)).Result()
	vs, _ := m.redisC.HGetAll(m.ctx, keys[0]).Result()
	return vs
}

// UserInfo for api
func (m *Manage) UserInfo(name string) (u *ent.User, e error) {
	if us, err := queryUsers(m.ctx, m.entC); err != nil {
		e = err
	} else {
		for _, eu := range us {
			if eu.Name == name {
				u = eu
			}
		}
	}
	return
}

// AddUser ...
func (m *Manage) AddUser(name, passwd, group string, adminFlag bool, phone ...string) error {
	eg, _ := queryGroupByName(m.ctx, m.entC, group)
	_, err := addUser(m.ctx, m.entC, name, passwd, eg, adminFlag, phone...)
	return err
}

// UpdateUser ...
func (m *Manage) UpdateUser(name, passwd, group string, adminFlag bool, phone ...string) error {
	eg, _ := queryGroupByName(m.ctx, m.entC, group)
	err := updateUser(m.ctx, m.entC, name, passwd, eg, adminFlag, phone...)
	return err
}

// UserAdminFlag ....
func (m *Manage) UserAdminFlag(uname string) bool {
	u, _ := queryUserByName(m.ctx, m.entC, uname)
	return userAdminFlag(m.ctx, u, uname)
}
