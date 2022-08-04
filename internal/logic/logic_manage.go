package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"github.com/yjiong/iotdor/ent"
	"github.com/yjiong/iotdor/internal/datasrc"
)

// Manage logic handle
type Manage struct {
	ctx context.Context
	datasrc.DSrcer
	entC     *ent.Client
	redisC   *redis.Client
	iotdName string
}

// MsgHandle ....
func (m *Manage) MsgHandle() {
	log.Infoln("start logic message handle")
	for {
		select {
		case <-m.ctx.Done():
			m.Close()
			m.entC.Close()
			m.redisC.Close()
			return
		default:
			msg := <-m.DataDownChan()
			if cmd, err := msg.Get("cmd").String(); err == nil {
				log.Debugf("receive mqtt data cmd=%s", cmd)
				gwID := msg.Get("sender").MustString()
				tstamp, _ := msg.Get("tstamp").Int64()
				tstr := time.Unix(tstamp, 0).Format("2006-01-02 15:04:05")
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
						if err = m.redisC.HSet(m.ctx, m.mkKeyPrefix(gwID, devID), md).Err(); err != nil {
							log.Error(err)
						}
					}
				}
			}
		}
	}
}

// NewManage ....
func NewManage(ctx context.Context, dsrc datasrc.DSrcer, entc *ent.Client, redisc *redis.Client, iname string) *Manage {
	return &Manage{
		ctx:      ctx,
		DSrcer:   dsrc,
		entC:     entc,
		redisC:   redisc,
		iotdName: iname,
	}
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
