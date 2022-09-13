package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	simplejson "github.com/bitly/go-simplejson"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// BASEPATH execfile path
var BASEPATH string
var cfgtype string
var version string

var dataDBpath string
var logLevel int

//var gorillaWS *api.WSClient

var rootCmd = &cobra.Command{
	Use:   "iotdor",
	Short: "LanChuang iotdor data collectien device project",
	Long: `	> documentation & support: http://www.xxx.com/
	> source & copyright information: https://github.com/yjiong/fkiotdor`,
	RunE: run,
}

func init() {
	eb, _ := exec.LookPath(os.Args[0])
	absp, _ := filepath.Abs(eb)
	BASEPATH = filepath.Dir(absp)
	if err := initDataSrcAndDBconfig(); err != nil {
		log.Fatalln(err)
	}
	rootCmd.PersistentFlags().IntVarP(&logLevel, "log-level", "L", 4, "debug=5, info=4, error=2, fatal=1, panic=0")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "config",
			Short: "Print the config",
			Run: func(cmd *cobra.Command, args []string) {
				prettyf(configViper.AllSettings())
			},
		},
	)
}

// Execute executes the root command.
func Execute(v string) {
	version = v
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func prettyf(v ...interface{}) {
	for _, s := range v {
		jb, _ := json.Marshal(s)
		sjb, _ := simplejson.NewJson(jb)
		pj, _ := sjb.EncodePretty()
		fmt.Println(string(pj))
	}
}

func initDataSrcAndDBconfig() error {
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
