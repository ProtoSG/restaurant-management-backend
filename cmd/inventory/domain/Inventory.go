package domain

import "restaurant-management-backend/cmd/inventory/domain/types"

type Inventory struct {
	Id             *types.InventoryId             `json:"id"`
	Name           *types.InventoryName           `json:"name"`
	ItemCategoryId *types.InventoryItemCategoryId `json:"item_category_id"`
	Quantity       *types.InventoryQuantity       `json:"quantity"`
	Price          *types.InventoryPrice          `json:"price"`
	Image          *types.InventoryImage          `json:"image"`
}

func NewInventory(id *types.InventoryId, name *types.InventoryName, itemCategoryId *types.InventoryItemCategoryId, quantity *types.InventoryQuantity, price *types.InventoryPrice, image *types.InventoryImage) *Inventory {
	return &Inventory{
		Id:             id,
		Name:           name,
		ItemCategoryId: itemCategoryId,
		Quantity:       quantity,
		Price:          price,
		Image:          image,
	}
}
