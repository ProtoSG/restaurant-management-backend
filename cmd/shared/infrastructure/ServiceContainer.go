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

	return container
}
