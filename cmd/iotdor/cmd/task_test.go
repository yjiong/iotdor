package cmd

import (
	"testing"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_task(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "01-02 15:04:05",
		FullTimestamp:   true,
	})
	Convey("==================test task=====================\n", t, func() {
		initDataSrcAndDB()
	})
}
