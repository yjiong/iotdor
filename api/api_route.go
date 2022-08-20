package api

import "github.com/gorilla/mux"

func configRoute(mr *mux.Router, dtr *IotdorTran) {
	mr.Use(dtr.validateToken, dtr.validateAdmin)
	mr.PathPrefix("/sys/user_info").HandlerFunc(dtr.userInfo)
	mr.PathPrefix("/sys/add_user").HandlerFunc(dtr.addUpdateUser).Methods("PUT")
	mr.PathPrefix("/sys/update_user").HandlerFunc(dtr.addUpdateUser).Methods("POST")
	mr.PathPrefix("/sys/delete_user/{uname}").HandlerFunc(dtr.delUser).Methods("DELETE")
}

func gatewayRoute(mr *mux.Router, dtr *IotdorTran) {
	mr.PathPrefix("/list").HandlerFunc(dtr.iotdorList)
	mr.Use(dtr.validateToken)
}

func deviceRoute(mr *mux.Router, dtr *IotdorTran) {
	mr.Use(dtr.validateToken)
	mr.PathPrefix("/realtime_data/{devid}").HandlerFunc(dtr.deviceRealTimeValue)
	mr.PathPrefix("/get_all_device_id").HandlerFunc(dtr.allDevices)
}
