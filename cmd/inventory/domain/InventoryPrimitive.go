package domain

import "restaurant-management-backend/cmd/shared/domain"

type InventoryPrimitive struct {
	Id             int     `json:"id"`
	Name           string  `json:"name"`
	ItemCategoryId int     `json:"itemCategoryId"`
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
		return &domain.ValidationFieldError{Field: "name", Message: "Name is required"}
	}
	if this.ItemCategoryId == 0 {
		return &domain.ValidationFieldError{Field: "itemCategoryId", Message: "ItemCategoryId is required"}
	}
	if this.Quantity == 0 {
		return &domain.ValidationFieldError{Field: "quantity", Message: "Quantity is required"}
	}
	if this.Price == 0 {
		return &domain.ValidationFieldError{Field: "price", Message: "Price is required"}
	}
	return nil
}
