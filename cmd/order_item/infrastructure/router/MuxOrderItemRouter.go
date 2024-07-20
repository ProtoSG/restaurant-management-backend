package router

import (
	"restaurant-management-backend/cmd/order_item/infrastructure/controller"
	"restaurant-management-backend/cmd/shared/infrastructure"

	"github.com/gorilla/mux"
)

func MuxOrderItemRouter(r *mux.Router, serviceContainer *infrastructure.ServiceContainer) {
	orderItemController := controller.NewHttpOrderItemController(serviceContainer)

	r.HandleFunc("/order_item", orderItemController.Create).Methods("POST")
	r.HandleFunc("/order_item", orderItemController.GetAll).Methods("GET")
	r.HandleFunc("/order_item/{id}", orderItemController.GetById).Methods("GET")
	r.HandleFunc("/order_item/{id}", orderItemController.Edit).Methods("PUT")
	r.HandleFunc("/order_item/{id}", orderItemController.Delete).Methods("DELETE")
}
