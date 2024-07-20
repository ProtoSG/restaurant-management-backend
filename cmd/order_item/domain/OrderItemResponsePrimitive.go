package domain

import domainInventory "restaurant-management-backend/cmd/inventory/domain"

type OrderItemResponsePrimitive struct {
	Id       int                                         `json:"id"`
	OrderId  int                                         `json:"order_id"`
	ItemId   int                                         `json:"item_id"`
	Item     *domainInventory.InventoryResponsePrimitive `json:"item"`
	Quantity int                                         `json:"quantity"`
	SubTotal float32                                     `json:"sub_total"`
}

func (this OrderItemResponse) MapToPrimitive() *OrderItemResponsePrimitive {
	return &OrderItemResponsePrimitive{
		Id:       this.Id.Value,
		OrderId:  this.OrderId.Value,
		ItemId:   this.ItemId.Value,
		Item:     this.Item.MapToPrimitive(),
		Quantity: this.Quantity.Value,
		SubTotal: this.SubTotal.Value,
	}
}
