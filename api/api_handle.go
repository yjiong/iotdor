package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/yjiong/iotdor/ent"
)

// ManageAPI ...
type ManageAPI interface {
	// user ....
	UsersInfo() ([]*ent.User, error)
	UserInfo(uName string) (*ent.User, error)
	AddUser(uName, passwd string, adminFlag bool, phone ...string) error
	UpdateUser(uName, passwd string, adminFlag bool, phone ...string) error
	UserAdminFlag(uName string) bool
	DelUser(uName string) error
	// gateway ....
	AllGateways() []*ent.Gateway
	// device ....
	AllDevices() []string
	DeviceRealTimeValue(devid string) map[string]string
	UpdateUserLoginInfo(uName, lip string) error
	// OrganizationTree ....
	OrganizationTree() ([]ent.OrganizationTree, error)
	AddOrganizationTree(o ent.OrganizationTree) error
	//UpdateOrganizationTree(oName string, o ent.OrganizationTree) error
	//DeleteOrganizationTree(oName string) error
	// OrganizationPosition ....
	//OrganizationPosition() ([]ent.OrganizationPosition, error)
	AddOrganizationPosition(o ent.OrganizationPosition) error
	UpdateOrganizationPosition(oName string, o ent.OrganizationPosition) error
	DeleteOrganizationPosition(oName string) error
	BeRelatedDeviceToOrganizationPosition(oName, devid string) error
	RemoveDeviceFromOrganizationPosition(devid string) error
	QueryOrganizationPositionDevices(oName string) ([]string, error)
}

func (dtr *IotdorTran) userInfo(w http.ResponseWriter, r *http.Request) {
	loc, _ := time.LoadLocation("Local")
	if usinfo, err := dtr.UsersInfo(); err == nil {
		usis := make([]map[string]any, 0)
		for _, uinfo := range usinfo {
			usis = append(usis, map[string]any{
				"name":            uinfo.Name,
				"phone":           uinfo.Phone,
				"last_login_time": uinfo.LastLoginTime.In(loc).Format("2006-01-02 15:04:05"),
				"last_login_ip":   uinfo.LastLoginIP,
				"admin_flag":      dtr.UserAdminFlag(uinfo.Name),
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

func (dtr *IotdorTran) organizationTree(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		var os []ent.OrganizationTree
		if os, err = dtr.OrganizationTree(); err == nil {
			rm := make([]map[string]interface{}, 0)
			for _, o := range os {
				rm = append(rm, map[string]interface{}{
					"ID":    o.ID,
					"Name":  o.Name,
					"Level": o.Level,
				})
			}
			respJSON(w, rm)
			return
		}
	case "POST":
		var ot ent.OrganizationTree
		decodeJSON(r, &ot)
		if err = dtr.AddOrganizationTree(ot); err == nil {
			respJSON(w, map[string]string{"msg": "create organizationtree successful"})
			return
		}
	case "PUT":
		respJSON(w, "TODO")
		return
	}
	//{
	////respJSON(w, dtr.)
	//} else {
	respError(200, w, err)
	//}
}
