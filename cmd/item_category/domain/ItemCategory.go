package domain

import "restaurant-management-backend/cmd/item_category/domain/types"

type ItemCategory struct {
	Id   *types.ItemCategoryId   `json:"id"`
	Name *types.ItemCategoryName `json:"name"`
}

func NewItemCategory(id *types.ItemCategoryId, name *types.ItemCategoryName) *ItemCategory {
	return &ItemCategory{
		Id:   id,
		Name: name,
	}
}
