package domain

import (
	typesInventory "restaurant-management-backend/cmd/inventory/domain/types"
	typesOrder "restaurant-management-backend/cmd/order/domain/types"
	"restaurant-management-backend/cmd/order_item/domain/types"
)

type OrderItem struct {
	Id       *types.OrderItemId          `json:"id"`
	OrderId  *typesOrder.OrderId         `json:"order_id"`
	ItemId   *typesInventory.InventoryId `json:"item_id"`
	Quantity *types.OrderItemQuantity    `json:"quantity"`
	SubTotal *types.OrderItemSubTotal    `json:"sub_total"`
}

func NewOrderItem(
	id *types.OrderItemId,
	orderId *typesOrder.OrderId,
	itemId *typesInventory.InventoryId,
	quantity *types.OrderItemQuantity,
	subTotal *types.OrderItemSubTotal,
) *OrderItem {
	return &OrderItem{
		Id:       id,
		OrderId:  orderId,
		ItemId:   itemId,
		Quantity: quantity,
		SubTotal: subTotal,
	}
}
