package router

import (
	"restaurant-management-backend/cmd/shared/infrastructure"
	"restaurant-management-backend/cmd/user/infrastructure/controller"

	"github.com/gorilla/mux"
)

func MuxUserRouter(r *mux.Router, serviceContainer *infrastructure.ServiceContainer) {
	userController := controller.NewHttpUserController(serviceContainer)

	r.HandleFunc("/users", userController.Create).Methods("POST")
	r.HandleFunc("/users", userController.GetAll).Methods("GET")
	r.HandleFunc("/users/{id}", userController.GetById).Methods("GET")
	r.HandleFunc("/users/{id}", userController.Edit).Methods("PUT")
	r.HandleFunc("/users/{id}", userController.Delete).Methods("DELETE")
}
