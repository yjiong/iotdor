package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	simplejson "github.com/bitly/go-simplejson"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	//"github.com/spf13/viper"
	//"github.com/yjiong/iotdor/api"
)

// BASEPATH execfile path
var BASEPATH string
var cfgtype string
var version string
var dataDBpath string
var logLevel int

//var gorillaWS *api.WSClient

var rootCmd = &cobra.Command{
	Use:   "lchj212",
	Short: "LanChuang Hj212 data collectien device project",
	Long: `	> documentation & support: http://www.xxx.com/
	> source & copyright information: https://github.com/yjiong/fkhj212`,
	RunE: run,
}

func init() {
	eb, _ := exec.LookPath(os.Args[0])
	absp, _ := filepath.Abs(eb)
	BASEPATH = filepath.Dir(absp)
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().IntVarP(&logLevel, "log-level", "L", 4, "debug=5, info=4, error=2, fatal=1, panic=0")
	rootCmd.PersistentFlags().StringVarP(&cfgtype, "configType", "t", "sqlite", "path to configuration file (optional)")
	//rootCmd.PersistentFlags().StringVarP(&dataDBpath, "databasePath", "d", filepath.Join(BASEPATH, "db"), "set database path")

	// bind flag to config vars
	//viper.BindPFlag("general.log_level", rootCmd.PersistentFlags().Lookup("log-level"))

	// defaults
	//viper.SetDefault("postgresql.automigrate", true)
	//viper.SetDefault("postgresql.dsn", "postgres://localhost/postgresql?sslmode=disable")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "config",
			Short: "Print the config",
			Run: func(cmd *cobra.Command, args []string) {
				prettyf("")
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

func initConfig() {
	if cfgtype == "json" {
		v := viper.New()
		v.SetConfigType("json")
		for f, s := range map[string]interface{}{
			"updataconf": "",
		} {
			v.SetConfigName(f)
			if err := v.ReadInConfig(); err != nil {
				switch err.(type) {
				case viper.ConfigFileNotFoundError:
					log.Warning("No configuration file found, using defaults.")
				default:
					log.WithError(err).Fatal("read configuration file error")
				}
			}
			if err := v.Unmarshal(s); err != nil {
				log.WithError(err).Fatal("unmarshal config error")
			}
			//log.Infoln(s)
		}
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
