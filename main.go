package main

import (
	"log"
	"net/http"
	itemCategory "restaurant-management-backend/cmd/item_category/infrastructure/router"
	"restaurant-management-backend/cmd/shared/infrastructure"
	tableRouter "restaurant-management-backend/cmd/table/infrastructure/router"
	tableCategoryRouter "restaurant-management-backend/cmd/table_category/infrastructure/router"
	userRouter "restaurant-management-backend/cmd/user/infrastructure/router"

	"github.com/gorilla/mux"
)

func main() {
	serviceContainer := infrastructure.NewServiceContainer()

	r := mux.NewRouter()

	userRouter.MuxUserRouter(r, serviceContainer)
	tableCategoryRouter.MuxTableCategoryRouter(r, serviceContainer)
	tableRouter.MuxTableRouter(r, serviceContainer)
	itemCategory.MuxItemCategoryRouter(r, serviceContainer)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
