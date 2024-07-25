package domain

import (
	"restaurant-management-backend/cmd/inventory/domain"
	typesInventory "restaurant-management-backend/cmd/inventory/domain/types"
	typesOrder "restaurant-management-backend/cmd/order/domain/types"
	"restaurant-management-backend/cmd/order_item/domain/types"
)

type OrderItemResponse struct {
	Id          *types.OrderItemId          `json:"id"`
	OrderId     *typesOrder.OrderId         `json:"order_id"`
	ItemId      *typesInventory.InventoryId `json:"item_id"`
	Item        *domain.InventoryResponse   `json:"item"`
	Quantity    *types.OrderItemQuantity    `json:"quantity"`
	SubTotal    *types.OrderItemSubTotal    `json:"sub_total"`
	Description *types.OrderItemDescription `json:"description"`
	Takeaway    *types.OrderItemTakeaway    `json:"takeaway"`
}

func NewOrderItemResponse(
	id *types.OrderItemId,
	orderId *typesOrder.OrderId,
	itemId *typesInventory.InventoryId,
	item *domain.InventoryResponse,
	quantity *types.OrderItemQuantity,
	subTotal *types.OrderItemSubTotal,
	description *types.OrderItemDescription,
	takeaway *types.OrderItemTakeaway,
) *OrderItemResponse {
	return &OrderItemResponse{
		Id:          id,
		OrderId:     orderId,
		ItemId:      itemId,
		Item:        item,
		Quantity:    quantity,
		SubTotal:    subTotal,
		Description: description,
		Takeaway:    takeaway,
	}
}
