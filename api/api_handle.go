package api

import (
	"fmt"
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
	UpdateUser(name, passwd string, adminFlag bool, phone ...string) error
	UserAdminFlag(uname string) bool
	DelUser(name string) error
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

func (dtr *IotdorTran) addUpdateUser(w http.ResponseWriter, r *http.Request) {
	type User struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		Phone     string `json:"phone"`
		AdminFlag bool   `json:"admin_flag"`
	}
	var u User
	decodeJSON(r, &u)
	var err error
	var parse string
	if r.Method == "POST" {
		parse = "add"
		err = dtr.AddUser(u.Username, u.Password, u.AdminFlag, u.Phone)
	} else if r.Method == "PUT" {
		parse = "update"
		err = dtr.UpdateUser(u.Username, u.Password, u.AdminFlag, u.Phone)
	}
	if err == nil {
		respJSON(w, map[string]string{"msg": fmt.Sprintf("%s user successful", parse)})
	} else {
		respError(200, w, err)
	}
}

func (dtr *IotdorTran) delUser(w http.ResponseWriter, r *http.Request) {
	if uname, ok := mux.Vars(r)["uname"]; ok {
		err := dtr.DelUser(uname)
		if err == nil {
			respJSON(w, map[string]string{"msg": "delete user successful"})
		} else {
			respError(200, w, err)
		}
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
