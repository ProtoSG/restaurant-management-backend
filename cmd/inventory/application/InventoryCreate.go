package application

import (
	"restaurant-management-backend/cmd/inventory/domain"
	"restaurant-management-backend/cmd/inventory/domain/repository"
	"restaurant-management-backend/cmd/inventory/domain/types"
)

type InventoryCreate struct {
	repository repository.InventoryRepository
}

func NewInventoryCreate(repository repository.InventoryRepository) *InventoryCreate {
	return &InventoryCreate{
		repository: repository,
	}
}

func (this InventoryCreate) Execute(id int, name string, item_category_id int, quantity int, price float32) error {
	inventoryId, err := types.NewInventoryId(id)
	if err != nil {
		return err
	}

	inventoryName, err := types.NewInventoryName(name)
	if err != nil {
		return err
	}

	inventoryItemCategoryId, err := types.NewInventoryItemCategoryId(item_category_id)
	if err != nil {
		return err
	}

	inventoryQuantity, err := types.NewInventoryQuantity(quantity)
	if err != nil {
		return err
	}

	inventoryPrice, err := types.NewInventoryPrice(price)
	if err != nil {
		return err
	}

	inventory := domain.NewInventory(inventoryId, inventoryName, inventoryItemCategoryId, inventoryQuantity, inventoryPrice)
	return this.repository.Create(inventory)
}
