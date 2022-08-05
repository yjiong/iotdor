package utils

import (
	"testing"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDirve(t *testing.T) {
	Convey("==================utils=====================\n", t, func() {
		log.SetLevel(log.DebugLevel)
		r, e := RunCmd("date")
		t.Log(r, e)
	})
}
