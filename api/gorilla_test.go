package api

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestWS(t *testing.T) {
	rand.Seed(time.Now().Unix())
	r := rand.Intn(0xFFFFFFFFFFFFFF)
	Convey("==================gorilla websocket=====================\n", t, func() {
		log.SetLevel(log.DebugLevel)
		ws := InitWSClint(fmt.Sprintf("2021%012X", r), "ws://xxxxxxxx/api/ws/transport")
		t.Log(ws)
		//ws.SendMessage(TranData{
		//UUID:            "xxxxxxxxxx",
		//Type:            "",
		//Data:            nil,
		//OriginalMessage: []byte{},
		//})
		<-time.After(time.Minute * 10)
	})
}
