package application

import (
	"restaurant-management-backend/cmd/order/domain"
	"restaurant-management-backend/cmd/order/domain/repository"
)

type OrderGetAll struct {
	repository repository.OrderRepository
}

func NewOrderGetAll(repository repository.OrderRepository) *OrderGetAll {
	return &OrderGetAll{repository: repository}
}

func (this OrderGetAll) Execute() ([]*domain.OrderResponse, error) {
	return this.repository.GetAll()
}
