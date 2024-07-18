package domain

import (
	"restaurant-management-backend/cmd/order/domain/types"
	typesTable "restaurant-management-backend/cmd/table/domain/types"
)

type Order struct {
	Id        *types.OrderId        `json:"id"`
	TableId   *typesTable.TableId   `json:"tableId"`
	UserId    *types.OrderUserId    `json:"userId"`
	Total     *types.OrderTotal     `json:"total"`
	CreatedAt *types.OrderCreatedAt `json:"createdAt"`
	UpdatedAt *types.OrderUpdatedAt `json:"updatedAt"`
}

func NewOrder(id *types.OrderId, tableId *typesTable.TableId, userId *types.OrderUserId, total *types.OrderTotal, createdAt *types.OrderCreatedAt, updatedAt *types.OrderUpdatedAt) *Order {
	return &Order{
		Id:        id,
		TableId:   tableId,
		UserId:    userId,
		Total:     total,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
