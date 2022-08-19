package cmd

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/go-redis/redis/v8"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/yjiong/iotdor/api"
	"github.com/yjiong/iotdor/internal/datasrc"
	"github.com/yjiong/iotdor/internal/logic"
)

//go:embed config.yml
var configYml embed.FS
var lm *logic.Manage

//var rawDB *sql.DB
//var dataSrc datasrc.DSrcer
//var dbClient *ent.Client
//var redisClient *redis.Client
var configViper *viper.Viper

func setLogLevel() error {
	fp := filepath.Join(BASEPATH, "log/iotdor.log-%Y%m%d")
	fw, err := rotatelogs.New(
		fp,
		rotatelogs.WithLinkName(filepath.Join(BASEPATH, "log/iotdor.log")),
		//rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
		rotatelogs.WithRotationSize(50*1024*1024),
		rotatelogs.WithRotationCount(5),
	)
	if err != nil {
		log.Fatalf("log file init failed")
	}
	log.SetOutput(io.MultiWriter(os.Stdout, fw))
	log.SetLevel(log.Level(logLevel))
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006/01/02 15:04:05",
	})
	return nil
}

func printStartMessage() error {
	log.WithFields(log.Fields{
		"version": version,
		"docs":    "http://github.com/yjiong",
	}).Info("starting iotdor process")
	return nil
}

func initDataSrcAndDB() error {
	conyml := filepath.Join(BASEPATH, "config.yml")
	sysconyml := filepath.Join("/etc/iotdor", "config.yml")
	var df *os.File
	var err error
	if _, err = os.Stat(sysconyml); err != nil {
		if os.IsNotExist(err) {
			if _, err = os.Stat(conyml); err != nil {
				if os.IsNotExist(err) {
					if df, err = os.OpenFile(conyml, os.O_WRONLY|os.O_CREATE, 0666); err == nil {
						sexyml, _ := configYml.Open("config.yml")
						io.Copy(df, sexyml)
						df.Sync()
					}
				}
			}
		}
	}
	configViper = viper.New()
	configViper.AddConfigPath("/etc/iotdor")
	configViper.AddConfigPath(BASEPATH)
	configViper.SetConfigName("config.yml")
	configViper.SetConfigType("yml")
	err = configViper.ReadInConfig()
	return err
}

func runLogicMsgHandle() error {
	b := configViper.GetStringMapString("broker")
	log.Debugln(b)
	dataSrc := datasrc.NewMQTTDSrcer(b)
	d := configViper.GetStringMapString("database")
	log.Debugln(d)
	rawDB := logic.OpenRawDB(d["type"], mkDBDns(d))
	dbClient := logic.OpenMigrate(d["type"], mkDBDns(d))
	r := configViper.GetStringMapString("redis")
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r["host"], r["port"]),
		Password: r["password"],
		DB:       0, // use default DB
	})
	if err := redisClient.Ping(ctx).Err(); err != nil {
		return err
	}
	lm = logic.NewManage(ctx,
		rawDB,
		dataSrc,
		dbClient,
		redisClient,
		configViper.GetString("broker.iotdname"),
		configViper.GetInt64("database.storage_interval"))
	go lm.MsgHandle()
	return nil
}

func startAPI() error {
	hs := configViper.GetStringMapString("http_server")
	if configViper.GetBool("http_server.debug_flag") {
		go api.APIserver(lm, hs["port"])
	}
	return nil
}

func mkDBDns(param map[string]string) (dns string) {
	//"mysql", "root:pass@tcp(localhost:3306)/test"
	//"postgres","host=<host> port=<port> user=<user> dbname=<database> password=<pass>"
	//"sqlite3", "file:ent?mode=memory&cache=shared&_fk=1"
	//"gremlin", "http://localhost:8182"
	//"postgres", "host=mqtt.yaojiong.top port=5432 user=iotd dbname=iotd password=yaojiong"
	switch param["type"] {
	case "mysql":
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
			param["user"],
			param["password"],
			param["host"],
			param["port"],
			param["dbname"])
	case "postgres":
		dns = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
			param["host"],
			param["port"],
			param["user"],
			param["dbname"],
			param["password"])
	}
	return
}
