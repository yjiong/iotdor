package logic

import (
	"context"

	"github.com/yjiong/iotdor/ent"
	"github.com/yjiong/iotdor/ent/device"
	"github.com/yjiong/iotdor/ent/gateway"
	"github.com/yjiong/iotdor/ent/group"
	"github.com/yjiong/iotdor/ent/migrate"
	"github.com/yjiong/iotdor/ent/user"

	log "github.com/sirupsen/logrus"

	// postgres driver
	_ "github.com/lib/pq"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// OpenMigrate ....
//"mysql", "root:pass@tcp(localhost:3306)/test"
//"postgres","host=<host> port=<port> user=<user> dbname=<database> password=<pass>"
//"sqlite3", "file:ent?mode=memory&cache=shared&_fk=1"
//"gremlin", "http://localhost:8182"

func OpenMigrate(driveName, dns string) *ent.Client {
	client, err := ent.Open(driveName, dns)
	if err != nil {
		log.Fatal(err)
	}
	if scerr := client.Schema.Create(context.Background(), migrate.WithDropColumn(true), migrate.WithDropIndex(true)); scerr != nil {
		log.Fatal(scerr)
	}
	log.Infof("%s client init ok", driveName)
	return client
}

/************************** crud group *********************/
func addGroup(ctx context.Context, c *ent.Client, gname string) (*ent.Group, error) {
	return c.Group.
		Create().
		SetName(gname).
		Save(ctx)
}

func queryGroups(ctx context.Context, c *ent.Client) ([]*ent.Group, error) {
	return c.Group.Query().All(ctx)
}

func queryGroupByName(ctx context.Context, c *ent.Client, gname string) (*ent.Group, error) {
	return c.Group.Query().Where(group.Name(gname)).Only(ctx)
}

func delGroupByID(ctx context.Context, c *ent.Client, id int) error {
	return c.Group.DeleteOneID(id).
		Exec(ctx)
}

/************************** crud group *********************/

/************************** crud user *********************/

func addUser(ctx context.Context,
	c *ent.Client,
	name, passwd string,
	group *ent.Group,
	isAdmin bool) (*ent.User, error) {
	tuc := c.User.Create().SetName(name).SetPasswd(passwd)
	if isAdmin {
		return tuc.AddManage(group).Save(ctx)
	}
	return tuc.AddGroups(group).Save(ctx)
}

func queryUsers(ctx context.Context, c *ent.Client) ([]*ent.User, error) {
	return c.User.Query().All(ctx)
}

func queryUserByName(ctx context.Context, c *ent.Client, name string) ([]*ent.User, error) {
	return c.User.Query().Where(user.Name(name)).All(ctx)
}

func delUserByID(ctx context.Context, c *ent.Client, id int) error {
	return c.User.DeleteOneID(id).
		Exec(ctx)
}

/************************** crud user *********************/

/************************** crud gateway *********************/
func addGateway(ctx context.Context,
	c *ent.Client,
	gwid, svrid, broker, installAt string,
	upInterval int) (*ent.Gateway, error) {
	return c.Gateway.Create().
		SetGwid(gwid).
		SetSvrid(svrid).
		SetBroker(broker).
		SetInstallationLocation(installAt).
		SetUpInterval(upInterval).
		Save(ctx)
}

func queryGateways(ctx context.Context, c *ent.Client) ([]*ent.Gateway, error) {
	return c.Gateway.Query().All(ctx)
}

func queryGatewayByGWID(ctx context.Context, c *ent.Client, gwid string) (*ent.Gateway, error) {
	return c.Gateway.Query().Where(gateway.Gwid(gwid)).Only(ctx)
}

func updateGateway(ctx context.Context,
	g *ent.Gateway,
	gwid, svrid, broker, installAt string,
	upInterval int) error {
	return g.Update().
		SetGwid(gwid).
		SetSvrid(svrid).
		SetBroker(broker).
		SetInstallationLocation(installAt).
		SetUpInterval(upInterval).
		Exec(ctx)
}

func setGatewayDelFlag(ctx context.Context,
	g *ent.Gateway,
	flag bool) error {
	return g.Update().
		SetDeleteFlag(flag).
		Exec(ctx)
}

func getDevicesByGWID(ctx context.Context, c *ent.Client, gwid string) ([]*ent.Device, error) {
	return c.Gateway.Query().Where(gateway.Gwid(gwid)).QueryDevices().All(ctx)
}

func delGatewayByGWID(ctx context.Context, c *ent.Client, gwid string) error {
	_, err := c.Gateway.Delete().Where((gateway.Gwid(gwid))).
		Exec(ctx)
	return err
}

func setGwGroup(ctx context.Context, gw *ent.Gateway, g *ent.Group) error {
	_, err := gw.Update().SetGroup(g).Save(ctx)
	//_, err := c.Gateway.UpdateOneID(gw.ID).SetGroup(g).Save(ctx)
	return err
}

/************************** crud gateway *********************/

/************************** crud device *********************/
func addDevice(ctx context.Context,
	c *ent.Client,
	did, dtype, daddr, conn, name string,
	gw *ent.Gateway) (*ent.Device, error) {
	return c.Device.Create().
		SetDevID(did).
		SetDevType(dtype).
		SetDevAddr(daddr).
		SetConn(conn).
		SetName(name).
		SetGateway(gw).
		Save(ctx)
}

func queryDevices(ctx context.Context, c *ent.Client) ([]*ent.Device, error) {
	return c.Device.Query().All(ctx)
}

func queryDeviceByDevID(ctx context.Context, c *ent.Client, devid string) (*ent.Device, error) {
	return c.Device.Query().Where(device.DevID(devid)).Only(ctx)
}

func getGatewayByDevID(ctx context.Context, c *ent.Client, devid string) (*ent.Gateway, error) {
	return c.Device.Query().Where(device.DevID(devid)).QueryGateway().Only(ctx)
}

func updateDevice(ctx context.Context, d *ent.Device,
	did, dtype, daddr, conn, name string,
	gw *ent.Gateway) error {
	return d.Update().
		SetDevID(did).
		SetDevType(dtype).
		SetDevAddr(daddr).
		SetConn(conn).
		SetName(name).
		SetGateway(gw).
		Exec(ctx)
}

func setDeviceDelFlag(ctx context.Context, d *ent.Device, flag bool) error {
	return d.Update().
		SetDeleteFlag(flag).
		Exec(ctx)
}

func delDeviceByID(ctx context.Context, c *ent.Client, id int) error {
	return c.Device.DeleteOneID(id).
		Exec(ctx)
}

/************************** crud device *********************/
