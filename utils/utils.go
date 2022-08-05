package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/go-cmd/cmd"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
	"time"
)

// UpdateSystemTime set system time
func UpdateSystemTime(dt string) (err error) {
	system := runtime.GOOS
	switch system {
	case "windows":
		{
			if _, err = RunCmd(`date ` + strings.Split(dt, " ")[0]); err == nil {
				_, err = RunCmd(`time ` + strings.Split(dt, " ")[1])
				go func() {
					time.After(time.Second * 3)
					RunCmd(`sc stop hj212 && sc start hj212`)
				}()
			}
		}
	case "linux":
		{
			if _, err = RunCmd(`date -s "` + dt + `"`); err == nil {
				_, err = RunCmd(`hwclock -w -u`)
				//校时间后cron job 失效
				go func() {
					time.After(time.Second * 3)
					RunCmd(`systemctl restart lchj212`)
				}()
			}
		}
	case "darwin":
		{
			_, err = RunCmd(`date -s  "` + dt + `"`)
		}
	}
	if err != nil {
		err = errors.Wrap(err, "set system time error!")
		log.Errorln(err)
	}
	return
}

func md5hex(data string) string {
	m := md5.New()
	m.Write([]byte(data))
	return hex.EncodeToString(m.Sum(nil))
}

// EqualValueToKeyValue .....
func EqualValueToKeyValue(ev string) (kv map[string]string) {
	kv = make(map[string]string)
	eqvs := strings.Split(ev, ",")
	for _, eqv := range eqvs {
		eq := strings.Split(eqv, "=")
		if len(eq) == 2 {
			kv[eq[0]] = eq[1]
		}
	}
	return
}

// Exists .....
func Exists(path string) bool {
	stat, err := os.Stat(path)
	return stat != nil && !os.IsNotExist(err)
}

// RunCmd ...
func RunCmd(name string, args ...string) ([]string, error) {
	cmd := cmd.NewCmd(name, args...)
	status := <-cmd.Start()
	return status.Stdout, status.Error
}