package repository

import (
	"restaurant-management-backend/cmd/order/domain"
	"restaurant-management-backend/cmd/order/domain/types"
)

type OrderRepository interface {
	Create(order *domain.Order) error
	GetAll() ([]*domain.OrderResponse, error)
	GetById(id *types.OrderId) (*domain.OrderResponse, error)
	Edit(order *domain.Order) error
	Delete(id *types.OrderId) error
}
