package application

import (
	"restaurant-management-backend/cmd/order/domain"
	"restaurant-management-backend/cmd/order/domain/repository"
	"restaurant-management-backend/cmd/order/domain/types"
)

type OrderGetById struct {
	repository repository.OrderRepository
}

func NewOrderGetById(repository repository.OrderRepository) *OrderGetById {
	return &OrderGetById{repository: repository}
}

func (this *OrderGetById) Execute(id int) (*domain.OrderResponse, error) {
	orderId, err := types.NewOrderId(id)
	if err != nil {
		return nil, err
	}

	order, _ := this.repository.GetById(orderId)
	if order == nil {
		return nil, *domain.NewOrderNotFound(orderId)
	}

	return order, nil
}
