package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/yjiong/iotdor/ent"
)

func (dtr *IotdorTran) getPositionDevices(posid string) []string {
	devs := make([]string, 0)
	pds, _ := dtr.ListOrganizationPositionDevices(posid)
	for _, p := range pds {
		devs = append(devs, p.DevID)
	}
	return devs
}

func (dtr *IotdorTran) getPositionUsers(posid string) []string {
	users := make([]string, 0)
	pcs, _ := dtr.ListOrganizationPositionCharges(posid)
	for _, p := range pcs {
		users = append(users, p.Name)
	}
	return users
}

func (dtr *IotdorTran) getPositionDevicesUsers(oes []*ent.OrganizationPosition) []map[string]any {
	rps := make([]map[string]any, 0)
	for _, o := range oes {
		rps = append(rps, map[string]any{
			"position_id":            o.PositionID,
			"address":                o.Address,
			"floor":                  o.Floor,
			"unit_no":                o.UnitNo,
			"longitude_and_latitude": o.LongitudeAndLatitude,
			"summary":                o.Summary,
			"devices":                dtr.getPositionDevices(o.PositionID),
			"users":                  dtr.getPositionUsers(o.PositionID),
		})
	}
	return rps
}

func (dtr *IotdorTran) organizationPositionDevices(w http.ResponseWriter, r *http.Request) {
	var err error
	var devs []*ent.Device
	if posid, ok := mux.Vars(r)["posid"]; ok {
		if devs, err = dtr.ListOrganizationPositionDevices(posid); err == nil {
			respJSON(w, devs)
			return
		}
	} else {
		err = errors.New("error: need posid")
	}
	respError(http.StatusOK, w, err)
}

func (dtr *IotdorTran) organizationPositionPersonCharge(w http.ResponseWriter, r *http.Request) {
	var err error
	if posid, ok := mux.Vars(r)["posid"]; ok {
		users := dtr.getPositionUsers(posid)
		respJSON(w, users)
		return
	}
	err = errors.New("error: need posid")
	respError(http.StatusOK, w, err)
}

func (dtr *IotdorTran) organizationPosition(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		ps, err := dtr.OrganizationPosition()
		if err == nil {
			respJSON(w, dtr.getPositionDevicesUsers(ps))
			return
		}
	case "POST":
		var op ent.OrganizationPosition
		decodeJSON(r, &op)
		if err = dtr.CreateOrganizationPosition(op); err == nil {
			respJSON(w, map[string]string{"msg": "create organizationPosition successful"})
			return
		}
	case "PUT":
		var op ent.OrganizationPosition
		decodeJSON(r, &op)
		if err = dtr.UpdateOrganizationPosition(op); err == nil {
			respJSON(w, map[string]string{"msg": "update OrganizationPosition successful"})
			return
		}
	}
	respError(http.StatusOK, w, err)
}

func (dtr *IotdorTran) organizationPositionDel(w http.ResponseWriter, r *http.Request) {
	var err error
	if posid, ok := mux.Vars(r)["posid"]; ok {
		if err = dtr.DeleteOrganizationPosition(posid); err == nil {
			respJSON(w, map[string]string{"msg": "delete OrganizationPosition successful"})
			return
		}
	} else {
		err = errors.New("error: need posid")
	}
	respError(http.StatusOK, w, err)
}

func (dtr *IotdorTran) personChargeOrganizationPosition(w http.ResponseWriter, r *http.Request) {
	var err error
	pid, iok := mux.Vars(r)["posid"]
	uname, piok := mux.Vars(r)["uname"]
	if iok && piok {
		if r.Method == "POST" {
			if err = dtr.SetPersonChargeWithOrganizationPosition(pid, uname); err == nil {
				respJSON(w, map[string]string{"msg": "set person in charge to OrganizationPosition successful"})
				return
			}
		} else {
			if err = dtr.RemovePersonChargeWithOrganizationPosition(pid, uname); err == nil {
				respJSON(w, map[string]string{"msg": "remove charge from organizationtree successful"})
				return
			}
		}
	} else {
		err = errors.New("error: need OrganizationPosition posid and and user name")
	}
	respError(http.StatusOK, w, err)
}

func (dtr *IotdorTran) addDeviceToOrganizationPosition(w http.ResponseWriter, r *http.Request) {
	var err error
	pid, iok := mux.Vars(r)["posid"]
	did, piok := mux.Vars(r)["devid"]
	if iok && piok {
		if r.Method == "POST" {
			if err = dtr.AddDeviceToOrganizationPosition(pid, did); err == nil {
				respJSON(w, map[string]string{"msg": "add device to OrganizationPosition successful"})
				return
			}
		} else {
			if err = dtr.RemoveDeviceFromOrganizationPosition(pid, did); err == nil {
				respJSON(w, map[string]string{"msg": "remove device from organizationtree successful"})
				return
			}
		}
	} else {
		err = errors.New("error: need OrganizationPosition posid and and device devid")
	}
	respError(http.StatusOK, w, err)
}
