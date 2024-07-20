package repository

import (
	typesOrder "restaurant-management-backend/cmd/order/domain/types"
	"restaurant-management-backend/cmd/order_item/domain"
	"restaurant-management-backend/cmd/order_item/domain/types"
)

type OrderItemRepository interface {
	Create(orderItem *domain.OrderItem) error
	GetAll() ([]*domain.OrderItemResponse, error)
	GetById(id *types.OrderItemId) (*domain.OrderItemResponse, error)
	Edit(orderItem *domain.OrderItem) error
	Delete(id *types.OrderItemId) error
	GetByOrder(id *typesOrder.OrderId) ([]*domain.OrderItemResponse, error)
}
