package cmd

import (
	"embed"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

//go:embed config.yml
var configYml embed.FS

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

func startGorilla() error {
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
	//v := viper.New()
	//v.AddConfigPath("/etc/iotdor")
	//v.AddConfigPath(BASEPATH)
	//v.SetConfigName("config.yml")
	//v.SetConfigType("yml")
	//err = v.ReadInConfig()
	c, err := os.OpenFile(conyml, os.O_RDONLY, 0666)
	b, err := ioutil.ReadAll(c)
	var f map[string]interface{}
	yaml.UnmarshalStrict(b, &f)
	log.Debugln(f)
	return err
}

func startAPI() error {
	//go api.StartHTTP(dataProcess)
	return nil
}
