package infrastructure

import (
	"database/sql"
	"log"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type ServiceContainer struct {
	User          UserService
	TableCategory TableCategoryService
	Table         TableService
	ItemCategory  ItemCategoryService
	Inventory     InventoryService
	Order         OrderService
	OrderItem     OrderItemService
}

func NewServiceContainer() *ServiceContainer {
	env := NewEnv()
	container := &ServiceContainer{}

	db, err := sql.Open("libsql", env.URL)
	if err != nil {
		log.Fatalf("failed to open db : %s", err)
	}

	container.User = NewUserService(db)
	container.TableCategory = NewTableCategoryService(db)
	container.Table = NewTableService(db)
	container.ItemCategory = NewItemCategoryService(db)
	container.Inventory = NewInventoryService(db)
	container.Order = NewOrderService(db)
	container.OrderItem = NewOrderItemService(db)

	return container
}
