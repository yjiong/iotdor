package logic

import (
	"context"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"github.com/yjiong/iotdor/ent"
	"github.com/yjiong/iotdor/internal/datasrc"
)

// Manage logic handle
type Manage struct {
	ctx context.Context
	datasrc.DSrcer
	entC   *ent.Client
	redisC *redis.Client
}

// MsgHandle ....
func (m *Manage) MsgHandle() {
	log.Infoln("start logic message handle")
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
			switch cmd {
			case AutoUpdata:
				if md, err := msg.Get("data").Map(); err == nil {
					if devID, ok := md["devid"].(string); ok {
						m.redisC.HSet(m.ctx, devID, md)
					}
				}
			}
		}
	}
}

// NewManage ....
func NewManage(ctx context.Context, dsrc datasrc.DSrcer, entc *ent.Client, redisc *redis.Client) *Manage {
	return &Manage{
		ctx:    ctx,
		DSrcer: dsrc,
		entC:   entc,
		redisC: redisc,
	}
}
