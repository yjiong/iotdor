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

	g, err := queryGroupByName(ctx, client, "Iotdor")
	t.Log(g, err)

	t.Log("-------addUser----------")
	u, err := addUser(ctx, client, "alice", "123456", g, false)
	t.Log(u, err)
	//t.Log("-------queryUser----------")
	//us, err := queryUserByName(ctx, client, "alice")
	//t.Log(us, err)
	//t.Log("-------delUser----------")
	//err = delUserByID(ctx, client, u.ID)
	//t.Log(err)
}
