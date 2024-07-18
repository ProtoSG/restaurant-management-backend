package application

import (
	"restaurant-management-backend/cmd/order/domain"
	"restaurant-management-backend/cmd/order/domain/repository"
	"restaurant-management-backend/cmd/order/domain/types"
)

type OrderDelete struct {
	repository repository.OrderRepository
}

func NewOrderDelete(repository repository.OrderRepository) *OrderDelete {
	return &OrderDelete{
		repository: repository,
	}
}

func (this *OrderDelete) Execute(id int) error {
	orderId, err := types.NewOrderId(id)
	if err != nil {
		return err
	}

	if order, _ := this.repository.GetById(orderId); order == nil {
		return domain.NewOrderNotFound(orderId)
	}

	return this.repository.Delete(orderId)
}
