package application

import (
	"restaurant-management-backend/cmd/order/domain"
	"restaurant-management-backend/cmd/order/domain/repository"
	"restaurant-management-backend/cmd/order/domain/types"
	"time"
)

type OrderCreate struct {
	repository repository.OrderRepository
}

func NewOrderCreate(repository repository.OrderRepository) *OrderCreate {
	return &OrderCreate{
		repository: repository,
	}
}

func (this *OrderCreate) Execute(id int, tableId int, userId int, total float32, createdAt time.Time, updatedAt time.Time) error {
	orderId, err := types.NewOrderId(id)
	if err != nil {
		return err
	}

	orderTableId, erro := types.NewOrderTableId(tableId)
	if erro != nil {
		return erro
	}

	orderUserId, err := types.NewOrderUserId(userId)
	if err != nil {
		return err
	}

	orderTotal, err := types.NewOrderTotal(total)
	if err != nil {
		return err
	}

	orderCreatedAt, err := types.NewOrderCreatedAt(createdAt)
	if err != nil {
		return err
	}

	orderUpdatedAt, err := types.NewOrderUpdatedAt(updatedAt)
	if err != nil {
		return err
	}

	order := domain.NewOrder(orderId, orderTableId, orderUserId, orderTotal, orderCreatedAt, orderUpdatedAt)
	return this.repository.Create(order)
}
