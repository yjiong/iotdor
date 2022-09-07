package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/yjiong/iotdor/ent"
)

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
			respJSON(w, map[string]string{"msg": "create organizationTree successful"})
			return
		}
	case "PUT":
		var ot ent.OrganizationTree
		decodeJSON(r, &ot)
		if err = dtr.UpdateOrganizationTree(ot); err == nil {
			respJSON(w, map[string]string{"msg": "update organizationTree successful"})
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
		respJSON(w, map[string]string{"msg": "add organizationTree successful"})
		return
	}
	respError(http.StatusOK, w, err)
}

func (dtr *IotdorTran) organizationTreeDel(w http.ResponseWriter, r *http.Request) {
	var err error
	if ids, ok := mux.Vars(r)["id"]; ok {
		idint, _ := strconv.Atoi(ids)
		if err = dtr.DeleteOrganizationTree(idint); err == nil {
			respJSON(w, map[string]string{"msg": "delete organizationTree successful"})
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
				respJSON(w, map[string]string{"msg": "relate positoin to organizationTree successful"})
				return
			}
		} else {
			if err = dtr.RemovePositioinFromOranizationTree(idint, pidint); err == nil {
				respJSON(w, map[string]string{"msg": "remove positoin from organizationTree successful"})
				return
			}
		}
	} else {
		err = errors.New("error: need organizationTree id and postion posid")
	}
	respError(http.StatusOK, w, err)
}
