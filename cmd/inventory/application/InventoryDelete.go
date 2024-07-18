package application

import (
	"restaurant-management-backend/cmd/inventory/domain"
	"restaurant-management-backend/cmd/inventory/domain/repository"
	"restaurant-management-backend/cmd/inventory/domain/types"
)

type InventoryDelete struct {
	repository repository.InventoryRepository
}

func NewInventoryDelete(repository repository.InventoryRepository) *InventoryDelete {
	return &InventoryDelete{
		repository: repository,
	}
}

func (this *InventoryDelete) Execute(id int) error {
	inventoryId, err := types.NewInventoryId(id)
	if err != nil {
		return err
	}

	if inventory, _ := this.repository.GetById(inventoryId); inventory == nil {
		return domain.NewInventoryNotFound(inventoryId)
	}

	return this.repository.Delete(inventoryId)
}
