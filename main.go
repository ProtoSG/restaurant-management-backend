package main

import (
	"log"
	"net/http"
	inventoryRouter "restaurant-management-backend/cmd/inventory/infrastructure/router"
	itemCategory "restaurant-management-backend/cmd/item_category/infrastructure/router"
	orderRouter "restaurant-management-backend/cmd/order/infrastructure/router"
	orderItemRouter "restaurant-management-backend/cmd/order_item/infrastructure/router"
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
	inventoryRouter.MuxInventoryRouter(r, serviceContainer)
	orderRouter.MuxOrderRouter(r, serviceContainer)
	orderItemRouter.MuxOrderItemRouter(r, serviceContainer)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
