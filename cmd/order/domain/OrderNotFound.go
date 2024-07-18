package domain

import (
	"fmt"
	"restaurant-management-backend/cmd/order/domain/types"
)

type OrderNotFound struct {
	OrderId *types.OrderId
}

func NewOrderNotFound(orderId *types.OrderId) *OrderNotFound {
	return &OrderNotFound{
		OrderId: orderId,
	}
}

func (this OrderNotFound) Error() string {
	return fmt.Sprintf("Order with id %d not found", this.OrderId)
}
