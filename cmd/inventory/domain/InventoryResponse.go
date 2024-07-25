package domain

import (
	"restaurant-management-backend/cmd/inventory/domain/types"
	"restaurant-management-backend/cmd/item_category/domain"
)

type InventoryResponse struct {
	Id       *types.InventoryId       `json:"id"`
	Name     *types.InventoryName     `json:"name"`
	Category *domain.ItemCategory     `json:"category"`
	Quantity *types.InventoryQuantity `json:"quantity"`
	Price    *types.InventoryPrice    `json:"price"`
	Image    *types.InventoryImage    `json:"image"`
}

func NewInventoryResponse(id *types.InventoryId, name *types.InventoryName, category *domain.ItemCategory, quantity *types.InventoryQuantity, price *types.InventoryPrice, image *types.InventoryImage) *InventoryResponse {
	return &InventoryResponse{
		Id:       id,
		Name:     name,
		Category: category,
		Quantity: quantity,
		Price:    price,
		Image:    image,
	}
}
