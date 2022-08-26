package logic

import (
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/yjiong/iotdor/ent/gateway"
	"github.com/yjiong/iotdor/utils"
)

var gwStat = []bool{false, true}

func (m *Manage) gatewayInfoHandle(gwID string, stat int) {
	if exist, _ := m.entC.Gateway.Query().Where(gateway.Gwid(gwID)).Exist(m.ctx); !exist {
		addGateway(m.ctx, m.entC, gwID, m.iotdName, "", "", gwStat[stat], 60)
	}
	if stat > 0 {
		if ret, err := m.mqttCmd(gwID, InitSysGet); err == nil {
			rm := ret.(map[string]interface{})
			burl := fmt.Sprintf("%s:%s", rm["mqtt_svr_addr"], rm["mqtt_svr_port"])
			ver := fmt.Sprintf("%s", rm["version"])
			uit, _ := strconv.Atoi(fmt.Sprintf("%s", rm["interval"]))
			updateGateway(m.ctx, m.entC, gwID, m.iotdName, burl, "", ver, gwStat[stat], uit)
		}
		if ret, err := m.mqttCmd(gwID, ListDevItems); err == nil {
			log.Debugln(ret)
			//TODO
		}
	} else {
		m.entC.Gateway.Update().Where(gateway.Gwid(gwID)).SetOnline(false).Exec(m.ctx)
	}
}

func (m *Manage) mqttCmd(gwID, cmdType string) (interface{}, error) {
	mid := utils.GetSnowflakeID()
	go m.SendData(fmt.Sprintf("%s/%s", gwID, m.iotdName), map[string]string{
		"cmd": cmdType,
		"seq": mid})
	return m.mia.StartAndWaitRet(mid, time.Second*5)
}
