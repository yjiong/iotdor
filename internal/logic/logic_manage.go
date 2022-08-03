package logic

import (
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"github.com/yjiong/iotdor/ent"
	"github.com/yjiong/iotdor/internal/datasrc"
)

// Manage logic handle
type Manage struct {
	datasrc.DSrcer
	entC   *ent.Client
	redisC *redis.Client
}

// MsgHandle ....
func (m *Manage) MsgHandle() {
	for msg := range m.DataDownChan() {
		if cmd, err := msg.Get("cmd").String(); err == nil {
			log.Debugf("receive mqtt data cmd=%s", cmd)
			switch cmd {
			case AutoUpdata:
				msg.Get("data")
			}
		}
	}
}
