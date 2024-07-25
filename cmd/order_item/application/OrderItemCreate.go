package application

import (
	typesInventory "restaurant-management-backend/cmd/inventory/domain/types"
	typesOrder "restaurant-management-backend/cmd/order/domain/types"
	"restaurant-management-backend/cmd/order_item/domain"
	"restaurant-management-backend/cmd/order_item/domain/repository"
	"restaurant-management-backend/cmd/order_item/domain/types"
)

type OrderItemCreate struct {
	repository repository.OrderItemRepository
}

func NewOrderItemCreate(repository repository.OrderItemRepository) *OrderItemCreate {
	return &OrderItemCreate{repository: repository}
}

func (this OrderItemCreate) Execute(id int, orderId int, itemId int, quantity int, subTotal float32, description string, takeaway int) error {
	orderItemId, err := types.NewOrderItemId(id)
	if err != nil {
		return err
	}

	orderItemOrderId, err := typesOrder.NewOrderId(orderId)
	if err != nil {
		return err
	}

	orderItemInventoryId, err := typesInventory.NewInventoryId(itemId)
	if err != nil {
		return err
	}

	orderItemQuantity, err := types.NewOrderQuantity(quantity)
	if err != nil {
		return err
	}

	orderItemSubTotal, err := types.NewOrderSubTotal(subTotal)
	if err != nil {
		return err
	}

	orderItemDescription, err := types.NewOrderItemDescription(description)
	if err != nil {
		return err
	}

	orderItemTakeaway, err := types.NewOrderItemTakeaway(takeaway)
	if err != nil {
		return err
	}

	orderItem := domain.NewOrderItem(orderItemId, orderItemOrderId, orderItemInventoryId, orderItemQuantity, orderItemSubTotal, orderItemDescription, orderItemTakeaway)

	return this.repository.Create(orderItem)
}
