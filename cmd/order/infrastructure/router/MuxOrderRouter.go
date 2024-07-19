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
	// r.HandleFunc("/order/{id}", orderController.).Methods("POST")
	// r.HandleFunc("/order/{id}", orderController.Create).Methods("POST")
	// r.HandleFunc("/order/{id}", orderController.Create).Methods("POST")
}
