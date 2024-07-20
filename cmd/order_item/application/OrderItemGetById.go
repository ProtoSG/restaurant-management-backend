package application

import (
	"restaurant-management-backend/cmd/order_item/domain"
	"restaurant-management-backend/cmd/order_item/domain/repository"
	"restaurant-management-backend/cmd/order_item/domain/types"
)

type OrderItemGetById struct {
	repository repository.OrderItemRepository
}

func NewOrderItemGetById(repository repository.OrderItemRepository) *OrderItemGetById {
	return &OrderItemGetById{repository: repository}
}

func (this OrderItemGetById) Execute(id int) (*domain.OrderItemResponse, error) {
	orderItemId, err := types.NewOrderItemId(id)
	if err != nil {
		return nil, err
	}

	orderItem, _ := this.repository.GetById(orderItemId)
	if orderItem == nil {
		return nil, domain.NewOrderNotFound(orderItemId)
	}

	return orderItem, nil
}
