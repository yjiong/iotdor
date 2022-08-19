package api

import "github.com/gorilla/mux"

func configRoute(mr *mux.Router, dtr *IotdorTran) {
	mr.Use(dtr.validateToken, dtr.validateAdmin)
}

func gatewayRoute(mr *mux.Router, dtr *IotdorTran) {
	mr.Use(dtr.validateToken)
	mr.PathPrefix("/list").HandlerFunc(dtr.iotdorList)
}

func deviceRoute(mr *mux.Router, dtr *IotdorTran) {
	mr.Use(dtr.validateToken)
}
