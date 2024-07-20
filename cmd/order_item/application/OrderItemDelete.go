package application

import (
	"restaurant-management-backend/cmd/order_item/domain"
	"restaurant-management-backend/cmd/order_item/domain/repository"
	"restaurant-management-backend/cmd/order_item/domain/types"
)

type OrderItemDelete struct {
	repository repository.OrderItemRepository
}

func NewOrderItemDeletee(repository repository.OrderItemRepository) *OrderItemDelete {
	return &OrderItemDelete{repository: repository}
}

func (this OrderItemDelete) Execute(id int) error {
	orderItemId, err := types.NewOrderItemId(id)
	if err != nil {
		return err
	}

	if orderItem, _ := this.repository.GetById(orderItemId); orderItem == nil {
		return domain.NewOrderNotFound(*&orderItemId)
	}

	return this.repository.Delete(orderItemId)
}
