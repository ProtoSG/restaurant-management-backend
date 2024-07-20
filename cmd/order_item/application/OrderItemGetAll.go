package application

import (
	"restaurant-management-backend/cmd/order_item/domain"
	"restaurant-management-backend/cmd/order_item/domain/repository"
)

type OrderItemGetAll struct {
	repository repository.OrderItemRepository
}

func NewOrderItemGetAll(repository repository.OrderItemRepository) *OrderItemGetAll {
	return &OrderItemGetAll{repository: repository}
}

func (this OrderItemGetAll) Execute() ([]*domain.OrderItemResponse, error) {
	return this.repository.GetAll()
}
