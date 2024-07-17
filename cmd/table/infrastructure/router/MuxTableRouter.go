package router

import (
	"restaurant-management-backend/cmd/shared/infrastructure"
	"restaurant-management-backend/cmd/table/infrastructure/controller"

	"github.com/gorilla/mux"
)

func MuxTableRouter(r *mux.Router, serviceContainer *infrastructure.ServiceContainer) {

	tableController := controller.NewHttpTableConstroller(serviceContainer)

	r.HandleFunc("/tables", tableController.Create).Methods("POST")
	r.HandleFunc("/tables", tableController.GetAll).Methods("GET")
	r.HandleFunc("/tables/{id}", tableController.GetById).Methods("GET")
	r.HandleFunc("/tables/{id}", tableController.Edit).Methods("PUT")
	r.HandleFunc("/tables/{id}", tableController.Delete).Methods("DELETE")
}
