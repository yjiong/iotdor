package logic

import (
	"context"
	"database/sql"
	"encoding/hex"

	esql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/yjiong/iotdor/ent"
	"github.com/yjiong/iotdor/ent/gateway"
	"github.com/yjiong/iotdor/ent/group"
	"github.com/yjiong/iotdor/ent/migrate"
	"github.com/yjiong/iotdor/ent/user"
	"golang.org/x/crypto/bcrypt"

	log "github.com/sirupsen/logrus"

	// postgres driver
	_ "github.com/lib/pq"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// OpenRawDB ,
func OpenRawDB(driveName, dns string) *sql.DB {
	rawdb, err := sql.Open(driveName, dns)
	if err != nil {
		log.Fatal(err)
	}
	log.Infoln(dns, ":connnect successful")
	return rawdb
}

// OpenMigrate ....
// "mysql", "root:pass@tcp(localhost:3306)/test?charset=utf8&parseTime=True"
// "postgres","host=<host> port=<port> user=<user> dbname=<database> password=<pass>"
// "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1"
// "gremlin", "http://localhost:8182"
func OpenMigrate(driverName, dns string) *ent.Client {
	client, err := ent.Open(driverName, dns)
	if err != nil {
		log.Fatal(err)
	}
	if scerr := client.Schema.Create(context.Background(), migrate.WithDropColumn(true), migrate.WithDropIndex(true)); scerr != nil {
		log.Fatal(scerr)
	}
	log.Infof("%s client init ok", driverName)
	// migrate mytables
	if drive, err := esql.Open(driverName, dns); err == nil {
		if m, err := schema.NewMigrate(drive, migrate.WithDropColumn(true), migrate.WithDropIndex(true)); err == nil {
			m.Create(context.Background(), myTables...)
			log.Infoln("create myTables ok !")
		}
		drive.Close()
	}
	log.Infof("ent client %s connected", dns)
	return client
}

func addGroupIfNotExist(ctx context.Context, c *ent.Client, gname string) (g *ent.Group, e error) {

	if g, e = c.Group.Query().Where(group.Name(gname)).Only(ctx); g == nil || e != nil {
		g, e = c.Group.Create().SetName(gname).Save(ctx)
		// add default user name=gname passwd=123456
		addUser(ctx, c, gname, "123456", g, true)
	}
	return
}

func generatePasswdString(p string) (string, error) {
	gp, err := bcrypt.GenerateFromPassword([]byte(p), 0)
	hps := hex.EncodeToString(gp)
	return hps, err
}

func addUser(ctx context.Context,
	c *ent.Client,
	name, passwd string,
	group *ent.Group,
	isAdmin bool,
	phone ...string) (*ent.User, error) {
	gpw, err := generatePasswdString(passwd)
	if err != nil {
		return nil, err
	}
	tuc := c.User.Create().SetName(name).SetPasswd(gpw).AddGroups(group)
	if len(phone) == 1 {
		tuc.SetPhone(phone[0])
	}
	if isAdmin {
		return tuc.AddAdmins(group).Save(ctx)
	}
	return tuc.AddGroups(group).Save(ctx)
}

func updateUser(ctx context.Context,
	c *ent.Client,
	name, passwd string,
	group *ent.Group,
	isAdmin bool,
	phone ...string) error {
	gpw, err := generatePasswdString(passwd)
	if err != nil {
		return err
	}
	tuc := c.User.Update().Where(user.Name(name)).SetPasswd(gpw)
	if len(phone) == 1 {
		tuc.SetPhone(phone[0])
	}
	if isAdmin {
		return tuc.AddAdmins(group).Exec(ctx)
	}
	return tuc.ClearAdmins().Exec(ctx)
}

func queryUsers(ctx context.Context, c *ent.Client) ([]*ent.User, error) {
	return c.User.Query().All(ctx)
}

func delUserByID(ctx context.Context, c *ent.Client, id int) error {
	return c.User.DeleteOneID(id).
		Exec(ctx)
}

func addGateway(ctx context.Context,
	c *ent.Client,
	gwid, svrid, broker, installAt, stat, ver string,
	upInterval int) (*ent.Gateway, error) {
	return c.Gateway.Create().
		SetGwid(gwid).
		SetSvrid(svrid).
		SetBroker(broker).
		SetInstallationLocation(installAt).
		SetStat(stat).
		SetVersion(ver).
		SetUpInterval(upInterval).
		Save(ctx)
}

func updateGateway(ctx context.Context,
	c *ent.Client,
	gwid, svrid, broker, installAt, ver, stat string,
	upInterval int) error {
	return c.Gateway.Update().Where(gateway.Gwid(gwid)).
		SetSvrid(svrid).
		SetBroker(broker).
		SetInstallationLocation(installAt).
		SetStat(stat).
		SetVersion(ver).
		SetUpInterval(upInterval).
		Exec(ctx)
}
