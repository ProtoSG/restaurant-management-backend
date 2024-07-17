package router

import (
	"restaurant-management-backend/cmd/shared/infrastructure"
	"restaurant-management-backend/cmd/table_category/infrastructure/controller"

	"github.com/gorilla/mux"
)

func MuxTableCategoryRouter(r *mux.Router, serviceContainer *infrastructure.ServiceContainer) {
	tableCategoryController := controller.NewHttpTableCategoryController(serviceContainer)

	r.HandleFunc("/table_category", tableCategoryController.Create).Methods("POST")
	r.HandleFunc("/table_category", tableCategoryController.GetAll).Methods("GET")
	r.HandleFunc("/table_category/{id}", tableCategoryController.GetById).Methods("GET")
	r.HandleFunc("/table_category/{id}", tableCategoryController.Edit).Methods("PUT")
	r.HandleFunc("/table_category/{id}", tableCategoryController.Delete).Methods("DELETE")
}
