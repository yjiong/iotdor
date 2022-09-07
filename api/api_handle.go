package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/yjiong/iotdor/ent"
	"github.com/yjiong/iotdor/utils"
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
		vs := dtr.DeviceRealTimeValue(devid)
		if len(vs) > 0 {
			respJSON(w, vs)
			return
		}
	}
	respError(http.StatusOK, w, errors.New("devid not exist"))
}

func (dtr *IotdorTran) deviceHistoryValue(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if devid, ok := params["devid"]; ok {
		var qs utils.QuerySection
		decodeJSON(r, &qs)
		respJSON(w, dtr.DeviceHistoryValue(devid, qs))
	} else {
		respError(http.StatusOK, w, errors.New("devid not exist"))
	}
}

func (dtr *IotdorTran) getTreePosition(oes []*ent.OrganizationTree) []map[string]any {
	rts := make([]map[string]any, 0)
	rtps := make([]map[string]any, 0)
	for _, o := range oes {
		ps, _ := dtr.ListOrganizationTreePositions(o)
		for _, p := range ps {
			rtps = append(rtps, map[string]any{
				"position_id":            p.PositionID,
				"address":                p.Address,
				"floor":                  p.Floor,
				"unit_no":                p.UnitNo,
				"longitude_and_latitude": p.LongitudeAndLatitude,
				"summary":                p.Summary,
			})
		}
		rts = append(rts, map[string]any{
			"id":        o.ID,
			"parent_id": o.ParentID,
			"name":      o.Name,
			"level":     o.Level,
			"positions": rtps,
		})
	}
	return rts
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
			respJSON(w, dtr.getTreePosition(os))
			return
		}
	case "POST":
		var ot ent.OrganizationTree
		decodeJSON(r, &ot)
		if err = dtr.CreateOrganizationSubNode(ot); err == nil {
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

func (dtr *IotdorTran) addOrganizationTree(w http.ResponseWriter, r *http.Request) {
	var err error
	var ot struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		LeftOrRight string `json:"left_or_right"`
	}
	decodeJSON(r, &ot)
	if err = dtr.AddOrganizationNode(ot.ID, ot.Name, ot.LeftOrRight); err == nil {
		respJSON(w, map[string]string{"msg": "add organizationtree successful"})
		return
	}
	respError(http.StatusOK, w, err)
}

func (dtr *IotdorTran) organizationTreeDel(w http.ResponseWriter, r *http.Request) {
	var err error
	if ids, ok := mux.Vars(r)["id"]; ok {
		idint, _ := strconv.Atoi(ids)
		if err = dtr.DeleteOrganizationTree(idint); err == nil {
			respJSON(w, map[string]string{"msg": "delete organizationtree successful"})
			return
		}
	} else {
		err = errors.New("error: need id")
	}
	respError(http.StatusOK, w, err)
}

func (dtr *IotdorTran) relatePositionToOrganizationTree(w http.ResponseWriter, r *http.Request) {
	var err error
	ids, iok := mux.Vars(r)["id"]
	pids, piok := mux.Vars(r)["posid"]
	if iok && piok {
		idint, _ := strconv.Atoi(ids)
		pidint, _ := strconv.Atoi(pids)
		if r.Method == "POST" {
			if err = dtr.RelatePositioinToOranizationTree(idint, pidint); err == nil {
				respJSON(w, map[string]string{"msg": "relate positoin to organizationtree successful"})
				return
			}
		} else {
			if err = dtr.RemovePositioinFromOranizationTree(idint, pidint); err == nil {
				respJSON(w, map[string]string{"msg": "remove positoin from organizationtree successful"})
				return
			}
		}
	} else {
		err = errors.New("error: need organizationtree id and postion posid")
	}
	respError(http.StatusOK, w, err)
}
