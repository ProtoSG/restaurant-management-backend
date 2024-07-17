package main

import (
	"log"
	"net/http"
	"restaurant-management-backend/cmd/shared/infrastructure"
	"restaurant-management-backend/cmd/user/infrastructure/controller"

	"github.com/gorilla/mux"
)

func main() {
	serviceContainer := infrastructure.NewServiceContainer()
	userController := controller.NewHttpUserController(serviceContainer)

	r := mux.NewRouter()

	r.HandleFunc("/users", userController.Create).Methods("POST")
	r.HandleFunc("/users", userController.GetAll).Methods("GET")
	r.HandleFunc("/users/{id}", userController.GetById).Methods("GET")
	r.HandleFunc("/users/{id}", userController.Edit).Methods("PUT")
	r.HandleFunc("/users/{id}", userController.Delete).Methods("DELETE")

	log.Println("Server running on port 8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
