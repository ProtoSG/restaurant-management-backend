package application

import (
	"restaurant-management-backend/cmd/order/domain"
	"restaurant-management-backend/cmd/order/domain/repository"
	"restaurant-management-backend/cmd/order/domain/types"
	"time"

	typesTable "restaurant-management-backend/cmd/table/domain/types"
	typesUser "restaurant-management-backend/cmd/user/domain/types"
)

type OrderEdit struct {
	repository repository.OrderRepository
}

func NewOrderEdit(repository repository.OrderRepository) *OrderEdit {
	return &OrderEdit{
		repository: repository,
	}
}

func (this *OrderEdit) Execute(id int, tableId int, userId int, total float32, createdAt time.Time, updatedAt time.Time, completed int) error {
	orderId, err := types.NewOrderId(id)
	if err != nil {
		return err
	}

	if order, _ := this.repository.GetById(orderId); order == nil {
		return domain.NewOrderNotFound(orderId)
	}

	orderTableId, err := typesTable.NewTableId(tableId)
	if err != nil {
		return err
	}

	orderUserId, err := typesUser.NewUserId(userId)
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

	orderCompleted, err := types.NewOrderCompleted(completed)
	if err != nil {
		return err
	}

	order := domain.NewOrder(orderId, orderTableId, orderUserId, orderTotal, orderCreatedAt, orderUpdatedAt, orderCompleted)
	return this.repository.Edit(order)
}
