package logic

import (
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/yjiong/iotdor/ent"
	"github.com/yjiong/iotdor/ent/gateway"
	"github.com/yjiong/iotdor/utils"
)

var gwStat = []string{"offline", "online"}

func (m *Manage) gatewayInfoHandle(gwID, gwCtag, ver string, stat int) {
	if egw, err := m.entC.Gateway.Query().Where(gateway.Gwid(gwID)).Only(m.ctx); egw == nil || err != nil {
		addGateway(m.ctx, m.entC, gwID, m.iotdName, "", "", gwStat[stat], ver, 60)
	}
	if stat > 0 {
		if ret, err := m.mqttCmd(gwID, InitSysGet); err == nil {
			rm := ret.(map[string]interface{})
			burl := fmt.Sprintf("%s:%s", rm["mqtt_svr_addr"], rm["mqtt_svr_port"])
			ver := fmt.Sprintf("%s", rm["version"])
			uit, _ := strconv.Atoi(fmt.Sprintf("%s", rm["interval"]))
			updateGateway(m.ctx, m.entC, gwID, m.iotdName, burl, "", ver, gwStat[stat], uit)
		}
		m.checkAndUpdateDev(gwID, gwCtag)
	} else {
		m.entC.Gateway.Update().Where(gateway.Gwid(gwID)).SetStat("offline").Exec(m.ctx)
	}
}

func (m *Manage) checkAndUpdateDev(gwID, gwCtag string) {
	if m.redisC.Get(m.ctx, m.mkRedisKeyPrefix(gwID, GatewayCtagValue)).String() != gwCtag {
		if egw, err := m.entC.Gateway.Query().Where(gateway.Gwid(gwID)).Only(m.ctx); egw != nil && err == nil {
			if err := m.redisC.Set(m.ctx, m.mkRedisKeyPrefix(gwID, GatewayCtagValue), gwCtag, 0).Err(); err != nil {
				log.Error(err)
			}
			if ret, err := m.mqttCmd(gwID, ListDevItems); err == nil {
				log.Debugln(ret)
				var eds []*ent.Device
				unMarshlDev(ret, &eds)
				log.Debugf("%+v", eds)
				for _, ed := range eds {
					m.addOrUpdateDevice(ed, egw)
				}
			}
		}
	}
}

func (m *Manage) mqttCmd(gwID, cmdType string) (any, error) {
	mid := utils.GetSnowflakeID()
	go m.SendData(fmt.Sprintf("%s/%s", gwID, m.iotdName), map[string]string{
		"cmd": cmdType,
		"seq": mid})
	return m.mia.StartAndWaitRet(mid, time.Second*5)
}
