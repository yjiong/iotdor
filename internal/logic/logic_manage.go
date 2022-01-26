package logic

import (
	"github.com/yjiong/iotdor/ent"
	"github.com/yjiong/iotdor/internal/datasrc"
)

// Manage logic handle
type Manage struct {
	datasrc.DSrcer
	*ent.Client
}

// MsgHandle ....
func (m *Manage) MsgHandle() {
	for msg := range m.DataDownChan() {
		cmd := msg.Get("cmd")
	}
}
