package router

import (
	"restaurant-management-backend/cmd/item_category/infrastructure/controller"
	"restaurant-management-backend/cmd/shared/infrastructure"

	"github.com/gorilla/mux"
)

func MuxItemCategoryRouter(r *mux.Router, serviceContainer *infrastructure.ServiceContainer) {
	itemCategoryController := controller.NewHttpItemCategoryController(serviceContainer)

	r.HandleFunc("/item_category", itemCategoryController.GetAll).Methods("GET")
	r.HandleFunc("/item_category", itemCategoryController.Create).Methods("POST")
	r.HandleFunc("/item_category/{id}", itemCategoryController.GetById).Methods("GET")
	r.HandleFunc("/item_category/{id}", itemCategoryController.Edit).Methods("PUT")
	r.HandleFunc("/item_category/{id}", itemCategoryController.Delete).Methods("DELETE")
}
