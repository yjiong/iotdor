package api

import "github.com/gorilla/mux"

func configRoute(mr *mux.Router, dtr *IotdorTran) {
	mr.Use(dtr.validateToken)
	mrSysRt := mr.PathPrefix("/sys").Subrouter()
	mrSysRt.Use(dtr.validateAdmin)
	mrSysRt.PathPrefix("/user_info").HandlerFunc(dtr.userInfo)
	mrSysRt.PathPrefix("/add_user").HandlerFunc(dtr.addUpdateUser).Methods("POST")
	mrSysRt.PathPrefix("/update_user").HandlerFunc(dtr.addUpdateUser).Methods("PUT")
	mrSysRt.PathPrefix("/delete_user/{uname}").HandlerFunc(dtr.delUser).Methods("DELETE")
	mrSysRt.PathPrefix("/organization_tree").HandlerFunc(dtr.organizationTree).Methods("GET", "POST", "PUT")
}

func gatewayRoute(mr *mux.Router, dtr *IotdorTran) {
	mr.Use(dtr.validateToken)
	mr.PathPrefix("/list").HandlerFunc(dtr.iotdorGWList)
}

func deviceRoute(mr *mux.Router, dtr *IotdorTran) {
	mr.Use(dtr.validateToken)
	mr.PathPrefix("/realtime_data/{devid}").HandlerFunc(dtr.deviceRealTimeValue)
	mr.PathPrefix("/get_all_device_id").HandlerFunc(dtr.allDevices)
}
