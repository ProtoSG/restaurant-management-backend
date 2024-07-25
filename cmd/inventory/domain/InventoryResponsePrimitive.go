package domain

import "restaurant-management-backend/cmd/item_category/domain"

type InventoryResponsePrimitive struct {
	Id       int                           `json:"id"`
	Name     string                        `json:"name"`
	Category *domain.ItemCategoryPrimitive `json:"category"`
	Quantity int                           `json:"quantity"`
	Price    float32                       `json:"price"`
	Image    string                        `json:"image"`
}

func (this InventoryResponse) MapToPrimitive() *InventoryResponsePrimitive {
	return &InventoryResponsePrimitive{
		Id:       this.Id.Value,
		Name:     this.Name.Value,
		Category: this.Category.MapToPrimitive(),
		Quantity: this.Quantity.Value,
		Price:    this.Price.Value,
		Image:    this.Image.Value,
	}
}
