package application

import (
	"restaurant-management-backend/cmd/inventory/domain"
	"restaurant-management-backend/cmd/inventory/domain/repository"
	"restaurant-management-backend/cmd/inventory/domain/types"
)

type InventoryEdit struct {
	repository repository.InventoryRepository
}

func NewInventoryEdit(repository repository.InventoryRepository) *InventoryEdit {
	return &InventoryEdit{
		repository: repository,
	}
}

func (this *InventoryEdit) Execute(id int, name string, category_id int, quantity int, price float32, image string) error {
	inventoryId, err := types.NewInventoryId(id)
	if err != nil {
		return err
	}

	if inventory, _ := this.repository.GetById(inventoryId); inventory == nil {
		return domain.NewInventoryNotFound(inventoryId)
	}

	inventoryName, err := types.NewInventoryName(name)
	if err != nil {
		return err
	}

	inventoryCategoryId, err := types.NewInventoryItemCategoryId(category_id)
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

	inventoryImage, err := types.NewInventoryImage(image)
	if err != nil {
		return err
	}

	inventory := domain.NewInventory(inventoryId, inventoryName, inventoryCategoryId, inventoryQuantity, inventoryPrice, inventoryImage)

	return this.repository.Edit(inventory)
}
