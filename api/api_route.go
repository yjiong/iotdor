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
	mrSysRt.PathPrefix("/organization_tree/{id}").HandlerFunc(dtr.organizationTree).Methods("GET")
	mrSysRt.PathPrefix("/organization_tree").HandlerFunc(dtr.organizationTree).Methods("GET", "POST", "PUT")
	mrSysRt.PathPrefix("/organization_tree/{id}").HandlerFunc(dtr.organizationTreeDel).Methods("DELETE")
	mrSysRt.PathPrefix("/add_organization_tree").HandlerFunc(dtr.addOrganizationTree).Methods("POST")
	mrSysRt.PathPrefix("/relate_position_to_organization_tree/{id}/{posid}").HandlerFunc(dtr.relatePositionToOrganizationTree).Methods("POST", "DELETE")

	mrSysRt.PathPrefix("/organization_position").HandlerFunc(dtr.organizationPosition).Methods("GET", "POST", "PUT")
	mrSysRt.PathPrefix("/organization_position/{posid}").HandlerFunc(dtr.organizationPositionDel).Methods("DELETE")
	mrSysRt.PathPrefix("/relate_position_to_organization_position/{id}/{posid}").HandlerFunc(dtr.addDeviceToOrganizationPosition).Methods("POST", "DELETE")
}

func gatewayRoute(mr *mux.Router, dtr *IotdorTran) {
	mr.Use(dtr.validateToken)
	mr.PathPrefix("/list").HandlerFunc(dtr.iotdorGWList)
}

func deviceRoute(mr *mux.Router, dtr *IotdorTran) {
	mr.Use(dtr.validateToken)
	mr.PathPrefix("/realtime_data/{devid}").HandlerFunc(dtr.deviceRealTimeValue)
	mr.PathPrefix("/history_data/{devid}").HandlerFunc(dtr.deviceHistoryValue)
	mr.PathPrefix("/get_all_device_id").HandlerFunc(dtr.allDevices)
}
