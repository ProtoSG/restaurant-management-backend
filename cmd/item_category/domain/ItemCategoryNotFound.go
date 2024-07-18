package domain

import (
	"fmt"
	"restaurant-management-backend/cmd/item_category/domain/types"
)

type ItemCategoryNotFound struct {
	ItemCategoryId *types.ItemCategoryId `json:"id"`
}

func NewItemCategoryNotFound(itemCategoryId *types.ItemCategoryId) *ItemCategoryNotFound {
	return &ItemCategoryNotFound{ItemCategoryId: itemCategoryId}
}

func (this *ItemCategoryNotFound) Error() string {
	return fmt.Sprintf("ItemCategory con el ID %d no encontrado", *this.ItemCategoryId)
}
