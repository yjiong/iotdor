package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/yjiong/iotdor/utils"
)

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
