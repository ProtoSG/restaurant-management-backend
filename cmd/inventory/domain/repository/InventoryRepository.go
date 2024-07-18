package repository

import (
	"restaurant-management-backend/cmd/inventory/domain"
	"restaurant-management-backend/cmd/inventory/domain/types"
)

type InventoryRepository interface {
	GetAll() ([]*domain.InventoryResponse, error)
	GetById(id *types.InventoryId) (*domain.InventoryResponse, error)
	Create(inventory *domain.Inventory) error
	Edit(inventory *domain.Inventory) error
	Delete(id *types.InventoryId) error
}
