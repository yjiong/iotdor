package logic

import (
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
func (m *Manage) addOrUpdateDevice(did, dtype, daddr, name string,
	ri, si int,
	conn map[string]any,
	gw *ent.Gateway) error {
	exist, _ := m.entC.Device.Query().Where(device.DevID(did)).Exist(m.ctx)
	if !exist {
		return m.entC.Device.Create().
			SetDevID(did).
			SetType(dtype).
			SetConn(conn).
			SetName(name).
			SetGateway(gw).
			SetReadInterval(ri).
			SetStoreInterval(si).
			Exec(m.ctx)
	}
	return m.entC.Device.Update().
		Where(device.DevID(did)).
		SetType(dtype).
		SetConn(conn).
		SetName(name).
		SetGateway(gw).
		SetReadInterval(ri).
		SetStoreInterval(si).
		Exec(m.ctx)
}
