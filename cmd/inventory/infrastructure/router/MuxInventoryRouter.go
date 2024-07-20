package router

import (
	"restaurant-management-backend/cmd/inventory/infrastructure/controller"
	"restaurant-management-backend/cmd/shared/infrastructure"

	"github.com/gorilla/mux"
)

func MuxInventoryRouter(r *mux.Router, serviceContainer *infrastructure.ServiceContainer) {
	inventoryController := controller.NewHttpInventoryController(serviceContainer)

	r.HandleFunc("/inventory", inventoryController.Create).Methods("POST")
	r.HandleFunc("/inventory", inventoryController.GetAll).Methods("GET")
	r.HandleFunc("/inventory/{id}", inventoryController.GetById).Methods("GET")
	r.HandleFunc("/inventory/{id}", inventoryController.Edit).Methods("PUT")
	r.HandleFunc("/inventory/{id}", inventoryController.Delete).Methods("DELETE")
}
