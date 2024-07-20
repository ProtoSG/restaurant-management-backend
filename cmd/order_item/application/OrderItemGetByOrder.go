package application

import (
	typesOrder "restaurant-management-backend/cmd/order/domain/types"
	"restaurant-management-backend/cmd/order_item/domain"
	"restaurant-management-backend/cmd/order_item/domain/repository"
)

type OrderItemGetByOrder struct {
	repository repository.OrderItemRepository
}

func NewOrderItemGetByOrder(repository repository.OrderItemRepository) *OrderItemGetByOrder {
	return &OrderItemGetByOrder{repository: repository}
}

func (this OrderItemGetByOrder) Execute(id int) ([]*domain.OrderItemResponse, error) {
	orderId, err := typesOrder.NewOrderId(id)
	if err != nil {
		return nil, err
	}

	return this.repository.GetByOrder(orderId)
}
