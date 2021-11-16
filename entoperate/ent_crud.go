package entoperate

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/yjiong/iotdor/ent"
	"github.com/yjiong/iotdor/ent/device"
	"github.com/yjiong/iotdor/ent/group"
	"github.com/yjiong/iotdor/ent/migrate"
	"github.com/yjiong/iotdor/ent/user"
	// postgres driver
	_ "github.com/lib/pq"
)

// OpenMigrate ....
//"mysql", "root:pass@tcp(localhost:3306)/test"
//"postgres","host=<host> port=<port> user=<user> dbname=<database> password=<pass>"
//"sqlite3", "file:ent?mode=memory&cache=shared&_fk=1"
//"gremlin", "http://localhost:8182"
//"postgres", "host=mqtt.yaojiong.top port=5432 user=iotd dbname=iotd password=yaojiong"
func OpenMigrate(driveName, dns string) *ent.Client {
	client, err := ent.Open(driveName, dns)
	if err != nil {
		log.Fatal(err)
	}
	if client.Schema.Create(context.Background(), migrate.WithDropColumn(true), migrate.WithDropIndex(true)) != nil {
		log.Fatal()
	}
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
/************************** crud device *********************/
func queryDevices(ctx context.Context, c *ent.Client) ([]*ent.Device, error) {
	return c.Device.Query().All(ctx)
}

func queryDeviceByID(ctx context.Context, c *ent.Client, devid string) (*ent.Device, error) {
	return c.Device.Query().Where(device.DevID(devid)).Only(ctx)
}

func delDeviceByID(ctx context.Context, c *ent.Client, id int) error {
	return c.Device.DeleteOneID(id).
		Exec(ctx)
}

/************************** crud device *********************/
