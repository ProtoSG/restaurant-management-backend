package infrastructure

import (
	"database/sql"
	"restaurant-management-backend/cmd/inventory/application"
	"restaurant-management-backend/cmd/inventory/infrastructure"
)

type InventoryService struct {
	Create  *application.InventoryCreate
	GetAll  *application.InventoryGetAll
	GetById *application.InventoryGetById
	Delete  *application.InventoryDelete
	Edit    *application.InventoryEdit
}

func NewInventoryService(db *sql.DB) InventoryService {
	inventoryContainer := infrastructure.NewSQLiteInventoryRepository(db)
	return InventoryService{
		Create:  application.NewInventoryCreate(inventoryContainer),
		GetAll:  application.NewInventoryGetAll(inventoryContainer),
		GetById: application.NewInventoryGetById(inventoryContainer),
		Delete:  application.NewInventoryDelete(inventoryContainer),
		Edit:    application.NewInventoryEdit(inventoryContainer),
	}
}
