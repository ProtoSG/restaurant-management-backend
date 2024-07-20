package domain

import (
	"fmt"
	"restaurant-management-backend/cmd/order_item/domain/types"
)

type OrderNotFound struct {
	OrderItemId *types.OrderItemId
}

func NewOrderNotFound(orderItemId *types.OrderItemId) *OrderNotFound {
	return &OrderNotFound{OrderItemId: orderItemId}
}

func (this OrderNotFound) Error() string {
	return fmt.Sprintf("OrderItem con ID %d no encontrado", this.OrderItemId)
}
