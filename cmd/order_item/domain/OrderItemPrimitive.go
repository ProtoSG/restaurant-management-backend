package domain

import "restaurant-management-backend/cmd/shared/domain"

type OrderItemPrimitive struct {
	Id       int     `json:"id"`
	OrderId  int     `json:"order_id"`
	ItemId   int     `json:"item_id"`
	Quantity int     `json:"quantity"`
	SubTotal float32 `json:"sub_total"`
}

func (this OrderItem) MapToPrimitive() *OrderItemPrimitive {
	return &OrderItemPrimitive{
		Id:       this.Id.Value,
		OrderId:  this.OrderId.Value,
		ItemId:   this.ItemId.Value,
		Quantity: this.Quantity.Value,
		SubTotal: this.SubTotal.Value,
	}
}

func (this OrderItemPrimitive) Validate() *domain.ValidationFieldError {
	if this.OrderId == 0 {
		return &domain.ValidationFieldError{Field: "order_id", Message: "order_id es requerido"}
	}
	if this.ItemId == 0 {
		return &domain.ValidationFieldError{Field: "item_id", Message: "item_id es requerido"}
	}
	if this.Quantity == 0 {
		return &domain.ValidationFieldError{Field: "quantity", Message: "quantity es requerido"}
	}
	if this.SubTotal == 0 {
		return &domain.ValidationFieldError{Field: "sub_total", Message: "sub_total es requerido"}
	}
	return nil
}
