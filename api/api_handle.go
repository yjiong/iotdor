package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/yjiong/iotdor/ent"
)

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
		respError(http.StatusOK, w, err)
	}
}

func (dtr *IotdorTran) delUser(w http.ResponseWriter, r *http.Request) {
	if uname, ok := mux.Vars(r)["uname"]; ok {
		err := dtr.DelUser(uname)
		if err == nil {
			respJSON(w, map[string]string{"msg": "delete user successful"})
		} else {
			respError(http.StatusOK, w, err)
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
		respError(http.StatusOK, w, errors.New("devid not exist"))
	}
}

func (dtr *IotdorTran) organizationTree(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		var os []*ent.OrganizationTree
		if ids, ok := mux.Vars(r)["id"]; ok {
			idint, _ := strconv.Atoi(ids)
			os, err = dtr.SubOrganizationFromID(idint)
		} else {
			os, err = dtr.OrganizationTree()
		}
		if err == nil {
			rm := make([]map[string]interface{}, 0)
			for _, o := range os {
				rm = append(rm, map[string]interface{}{
					"ID":       o.ID,
					"ParentID": o.ParentID,
					"Name":     o.Name,
					"Level":    o.Level,
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
		var ot ent.OrganizationTree
		decodeJSON(r, &ot)
		if err = dtr.UpdateOrganizationTree(ot); err == nil {
			respJSON(w, map[string]string{"msg": "update organizationtree successful"})
			return
		}
	}
	respError(http.StatusOK, w, err)
}
