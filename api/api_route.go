package api

import "github.com/gorilla/mux"

func configRoute(mr *mux.Router, dtr *IotdorTran) {
	mr.Use(dtr.validateToken)
	mrSysRt := mr.PathPrefix("/sys").Subrouter()
	mrSysRt.Use(dtr.validateAdmin)
	mrSysRt.HandleFunc("/user_info", dtr.userInfo)
	mrSysRt.HandleFunc("/add_user", dtr.addUpdateUser).Methods("POST")
	mrSysRt.HandleFunc("/update_user", dtr.addUpdateUser).Methods("PUT")
	mrSysRt.HandleFunc("/delete_user/{uname}", dtr.delUser).Methods("DELETE")
	mrSysRt.HandleFunc("/organization_tree/{id}", dtr.organizationTree).Methods("GET")
	mrSysRt.HandleFunc("/organization_tree", dtr.organizationTree).Methods("GET", "POST", "PUT")
	mrSysRt.HandleFunc("/organization_tree/{id}", dtr.organizationTreeDel).Methods("DELETE")
	mrSysRt.HandleFunc("/add_organization_tree", dtr.addOrganizationTree).Methods("POST")
	mrSysRt.HandleFunc("/relate_position_to_organization_tree/{id}/{posid}", dtr.relatePositionToOrganizationTree).Methods("POST", "DELETE")

	mrSysRt.HandleFunc("/organization_position", dtr.organizationPosition).Methods("GET", "POST", "PUT")
	mrSysRt.HandleFunc("/organization_position/{posid}", dtr.organizationPositionDel).Methods("DELETE")
	mrSysRt.HandleFunc("/organization_position_person_charge/{posid}/{uname}", dtr.personChargeOrganizationPosition).Methods("POST", "DELETE")
	mrSysRt.HandleFunc("/organization_position_and_device/{posid}/{devid}", dtr.addDeviceToOrganizationPosition).Methods("POST", "DELETE")
	mrSysRt.HandleFunc("/organization_position_device/{posid}", dtr.organizationPositionDevices).Methods("GET")
	mrSysRt.HandleFunc("/organization_position_person_charge/{posid}", dtr.organizationPositionPersonCharge).Methods("GET")
}

func gatewayRoute(mr *mux.Router, dtr *IotdorTran) {
	mr.Use(dtr.validateToken)
	mr.HandleFunc("/list", dtr.iotdorGWList)
}

func deviceRoute(mr *mux.Router, dtr *IotdorTran) {
	mr.Use(dtr.validateToken)
	mr.HandleFunc("/realtime_data/{devid}", dtr.deviceRealTimeValue)
	mr.HandleFunc("/history_data/{devid}", dtr.deviceHistoryValue)
	mr.HandleFunc("/get_all_device_id", dtr.allDevices)
	mr.HandleFunc("/get_device_info/{devid}", dtr.deviceInfo)
}
