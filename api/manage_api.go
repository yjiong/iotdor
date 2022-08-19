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
	UsersInfo() ([]*ent.User, error)
	AddUser(name, passwd string, adminFlag bool, phone ...string) error
	UpdateUser(name, passwd, group string, adminFlag bool, phone ...string) error
	UserAdminFlag(uname string) bool
	AllDevices() []string
	DeviceRealTimeValue(devid string) map[string]string
}

func (dtr *IotdorTran) userInfo(w http.ResponseWriter, r *http.Request) {
	if usinfo, err := dtr.UsersInfo(); err == nil {
		usis := make([]map[string]any, 0)
		for _, uinfo := range usinfo {
			usis = append(usis, map[string]any{
				"name":       uinfo.Name,
				"phone":      uinfo.Phone,
				"admin_flag": dtr.UserAdminFlag(uinfo.Name),
			})
		}
		respJSON(w, usis)
	}
}

func (dtr *IotdorTran) addUser(w http.ResponseWriter, r *http.Request) {
	type User struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		Phone     string `json:"phone"`
		AdminFlag bool   `json:"admin_flag"`
	}
	var u User
	decodeJSON(r, &u)
	if err := dtr.AddUser(u.Username, u.Password, u.AdminFlag, u.Phone); err == nil {
		respJSON(w, map[string]string{"msg": "add user successful"})
	} else {
		respError(200, w, err)
	}
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
