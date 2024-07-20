package application

import (
	typesInventory "restaurant-management-backend/cmd/inventory/domain/types"
	typesOrder "restaurant-management-backend/cmd/order/domain/types"
	"restaurant-management-backend/cmd/order_item/domain"
	"restaurant-management-backend/cmd/order_item/domain/repository"
	"restaurant-management-backend/cmd/order_item/domain/types"
)

type OrderItemEdit struct {
	repository repository.OrderItemRepository
}

func NewOrderItemEdit(repository repository.OrderItemRepository) *OrderItemEdit {
	return &OrderItemEdit{repository: repository}
}

func (this OrderItemEdit) Execute(id int, orderId int, itemId int, quantity int, subTotal float32) error {
	orderItemId, err := types.NewOrderItemId(id)
	if err != nil {
		return err
	}

	if orderItem, _ := this.repository.GetById(orderItemId); orderItem == nil {
		return domain.NewOrderNotFound(orderItemId)
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

	orderItem := domain.NewOrderItem(orderItemId, orderItemOrderId, orderItemInventoryId, orderItemQuantity, orderItemSubTotal)

	return this.repository.Edit(orderItem)
}
