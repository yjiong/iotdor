package entoperate

import (
	"context"
	"testing"
)

func TestENT(t *testing.T) {
	t.Log("++++++++++++++++++test ent++++++++++++++++++++")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client := OpenMigrate("postgres", "host=mqtt.yaojiong.top port=5432 user=iotd dbname=iotd password=yaojiong")
	defer client.Close()
	/*************
	t.Log("-------addGroup----------")
	g, err := addGroup(ctx, client, "Ixxxotdor")
	t.Log(g, err)
	t.Log("-------queryGroupByName----------")
	g, err = queryGroupByName(ctx, client, "Ixxxotdor")
	t.Log(g, err)
	t.Log("-------delGroupByName----------")
	err = delGroupByID(ctx, client, g.ID)
	t.Log(err)
	***************/

	//t.Log(g, err)

	//t.Log("-------addUser----------")
	//u, err := addUser(ctx, client, "alice", "123456", g, false)
	//t.Log(u, err)
	t.Log("-------queryUser----------")
	us, err := queryUserByName(ctx, client, "alice")
	t.Log(us, err)
	//t.Log("-------delUser----------")
	//err = delUserByID(ctx, client, u.ID)
	//t.Log(err)

	t.Log("-------addGateway----------")
	gw, err := addGateway(ctx, client, "gwtest-001", "mqtt.yaojiong.top:1883", "newyork", 10)
	t.Log(gw, err)
	t.Log("-------queryGateway----------")
	gws, err := queryGateways(ctx, client)
	t.Log(gws, err)
	t.Log("-------setBelong----------")
	setBelong(ctx, gw, us[0])
	//t.Log("-------delGateway----------")
	//err = delGatewayByGWID(ctx, client, "gwtest-001")
	//t.Log(err)

	//t.Log("-------addDevice----------")
	//d, err := addDevice(ctx)
	//t.Log(d, err)
	//t.Log("-------queryDevice----------")
	//us, err := queryDeviceByName(ctx, client, "alice")
	//t.Log(us, err)
	//t.Log("-------delDevice----------")
	//err = delUserByID(ctx, client, u.ID)
	//t.Log(err)
}
