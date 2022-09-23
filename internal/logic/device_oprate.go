package logic

import (
	"github.com/bitly/go-simplejson"
	log "github.com/sirupsen/logrus"
	"github.com/yjiong/iotdor/ent"
	"github.com/yjiong/iotdor/ent/device"
	"github.com/yjiong/iotdor/ent/gateway"
)

func (m *Manage) setDevDelFlag(devid string) error {
	return m.entC.Device.Update().Where(device.DevID(devid)).SetDeleteFlag(true).Exec(m.ctx)
}
func (m *Manage) setGwAllDevDelFlag(gwid string) {
	if ds, err := m.entC.Device.Query().Where(device.HasGatewayWith(gateway.Gwid(gwid))).All(m.ctx); err == nil {
		for _, d := range ds {
			d.Update().SetDeleteFlag(true).Exec(m.ctx)
		}
	}
}
func (m *Manage) addOrUpdateDevice(dev *ent.Device, gw *ent.Gateway) error {
	//if gw, err = m.entC.Gateway.Query().Where(gateway.Gwid(gwid)).Only(m.ctx); gw != nil && err == nil {
	exist, _ := m.entC.Device.Query().Where(device.DevID(dev.DevID)).Exist(m.ctx)
	if !exist {
		return m.entC.Device.Create().
			SetDevID(dev.DevID).
			SetType(dev.Type).
			SetConn(dev.Conn).
			SetName(dev.Name).
			SetGateway(gw).
			SetReadInterval(dev.ReadInterval).
			SetStoreInterval(dev.StoreInterval).
			SetDeleteFlag(false).
			Exec(m.ctx)
	}
	return m.entC.Device.Update().
		Where(device.DevID(dev.DevID)).
		SetType(dev.Type).
		SetConn(dev.Conn).
		SetName(dev.Name).
		SetGateway(gw).
		SetReadInterval(dev.ReadInterval).
		SetStoreInterval(dev.StoreInterval).
		SetDeleteFlag(false).
		Exec(m.ctx)
}

func unMarshlDev(aed any, ueds *[]ent.Device) {
	if sj, ok := aed.(*simplejson.Json); ok {
		log.Debugln(sj)
	}
	if eds, ok := aed.([]any); ok {
		for _, ed := range eds {
			unMarshlDev(ed, ueds)
		}
	} else {
		if aem, ok := aed.(map[string]any); ok {
			ed := ent.Device{
				DevID: aem["devid"].(string),
				Type:  aem["type"].(string),
				Conn:  aem["conn"].(map[string]any),
				//ReadInterval:  aem["read_interval"].(json.Number),
				//StoreInterval: aem["store_interval"].(json.Number),
				//Name:       aem["devname"].(string),
				DeleteFlag: false,
				//Summary:    aem["summary"].(string),
			}
			*ueds = append(*ueds, ed)
		}
	}
}
