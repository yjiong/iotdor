package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/yjiong/iotdor/ent"
)

func (dtr *IotdorTran) getPositionDevice(oes []*ent.OrganizationPosition) []map[string]any {
	rts := make([]map[string]any, 0)
	rtps := make([]map[string]any, 0)
	for _, o := range oes {
		ps, _ := dtr.ListOrganizationPositionDevices(o)
		for _, p := range ps {
			rtps = append(rtps, map[string]any{
				"devid": p.DevID,
			})
		}
		rts = append(rts, map[string]any{
			"position_id":            o.PositionID,
			"address":                o.Address,
			"floor":                  o.Floor,
			"unit_no":                o.UnitNo,
			"longitude_and_latitude": o.LongitudeAndLatitude,
			"summary":                o.Summary,
		})
	}
	return rts
}

func (dtr *IotdorTran) organizationPosition(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		ps, err := dtr.OrganizationPosition()
		if err == nil {
			respJSON(w, dtr.getPositionDevice(ps))
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
			if err = dtr.RemoveDeviceFromOrganizationPosition(did); err == nil {
				respJSON(w, map[string]string{"msg": "remove device from organizationtree successful"})
				return
			}
		}
	} else {
		err = errors.New("error: need OrganizationPosition posid and and device devid")
	}
	respError(http.StatusOK, w, err)
}
