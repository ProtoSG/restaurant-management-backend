package domain

import "restaurant-management-backend/cmd/shared/domain"

type InventoryPrimitive struct {
	Id             int     `json:"id"`
	Name           string  `json:"name"`
	ItemCategoryId int     `json:"item_category_id"`
	Quantity       int     `json:"quantity"`
	Price          float32 `json:"price"`
}

func (this Inventory) MapToPrimitive() *InventoryPrimitive {
	return &InventoryPrimitive{
		Id:             this.Id.Value,
		Name:           this.Name.Value,
		ItemCategoryId: this.ItemCategoryId.Value,
		Quantity:       this.Quantity.Value,
		Price:          this.Price.Value,
	}
}

func (this *InventoryPrimitive) Validate() *domain.ValidationFieldError {
	if this.Name == "" {
		return &domain.ValidationFieldError{Field: "name", Message: "name is required"}
	}
	if this.ItemCategoryId == 0 {
		return &domain.ValidationFieldError{Field: "item_category_id", Message: "item_category_id is required"}
	}
	if this.Quantity == 0 {
		return &domain.ValidationFieldError{Field: "quantity", Message: "quantity is required"}
	}
	if this.Price == 0 {
		return &domain.ValidationFieldError{Field: "price", Message: "price is required"}
	}
	return nil
}
