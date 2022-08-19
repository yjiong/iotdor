package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/yjiong/iotdor/ent"
)

// ManageAPI ...
type ManageAPI interface {
	UserInfo(name string) (*ent.User, error)
	AddUser(name, passwd, group string, adminFlag bool, phone ...string) error
	UpdateUser(name, passwd, group string, adminFlag bool, phone ...string) error
	UserAdminFlag(uname string) bool
	AllDevices() []string
	DeviceRealTimeValue(devid string) map[string]string
}

func (dtr *IotdorTran) allDevices(w http.ResponseWriter, r *http.Request) {
	respJSON(w, dtr.AllDevices())
}

func (dtr *IotdorTran) deviceRealTimeValue(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if devid, ok := params["devid"]; ok {
		respJSON(w, dtr.DeviceRealTimeValue(devid))
	} else {
		respError(200, w, errors.New("devid not exist"))
	}
}
