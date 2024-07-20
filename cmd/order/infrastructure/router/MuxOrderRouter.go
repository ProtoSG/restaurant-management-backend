package router

import (
	"restaurant-management-backend/cmd/order/infrastructure/controller"
	"restaurant-management-backend/cmd/shared/infrastructure"

	"github.com/gorilla/mux"
)

func MuxOrderRouter(r *mux.Router, serviceContainer *infrastructure.ServiceContainer) {
	orderController := controller.NewHttpOrderController(serviceContainer)

	r.HandleFunc("/order", orderController.Create).Methods("POST")
	r.HandleFunc("/order", orderController.GetAll).Methods("GET")
	r.HandleFunc("/order/{id}", orderController.GetById).Methods("GET")
	r.HandleFunc("/order/{id}", orderController.Edit).Methods("PUT")
	r.HandleFunc("/order/{id}", orderController.Delete).Methods("DELETE")
}
