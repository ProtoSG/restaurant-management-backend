package domain

import (
	"restaurant-management-backend/cmd/order/domain/types"
	typesTable "restaurant-management-backend/cmd/table/domain/types"
	typesUser "restaurant-management-backend/cmd/user/domain/types"
)

type Order struct {
	Id        *types.OrderId        `json:"id"`
	TableId   *typesTable.TableId   `json:"tableId"`
	UserId    *typesUser.UserId     `json:"userId"`
	Total     *types.OrderTotal     `json:"total"`
	CreatedAt *types.OrderCreatedAt `json:"createdAt"`
	UpdatedAt *types.OrderUpdatedAt `json:"updatedAt"`
	Completed *types.OrderCompleted `json:"completed"`
}

func NewOrder(
	id *types.OrderId,
	tableId *typesTable.TableId,
	userId *typesUser.UserId,
	total *types.OrderTotal,
	createdAt *types.OrderCreatedAt,
	updatedAt *types.OrderUpdatedAt,
	completed *types.OrderCompleted,
) *Order {
	return &Order{
		Id:        id,
		TableId:   tableId,
		UserId:    userId,
		Total:     total,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Completed: completed,
	}
}
