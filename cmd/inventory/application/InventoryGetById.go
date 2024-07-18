package application

import (
	"restaurant-management-backend/cmd/inventory/domain"
	"restaurant-management-backend/cmd/inventory/domain/repository"
	"restaurant-management-backend/cmd/inventory/domain/types"
)

type InventoryGetById struct {
	repository repository.InventoryRepository
}

func NewInventoryGetById(repository repository.InventoryRepository) *InventoryGetById {
	return &InventoryGetById{
		repository: repository,
	}
}

func (this *InventoryGetById) Execute(id int) (*domain.InventoryResponse, error) {
	inventoryId, err := types.NewInventoryId(id)
	if err != nil {
		return nil, err
	}

	inventory, err := this.repository.GetById(inventoryId)
	if err != nil {
		return nil, domain.NewInventoryNotFound(inventoryId)
	}

	return inventory, nil
}
