package domain

import (
	"restaurant-management-backend/cmd/order/domain/types"
	domainTable "restaurant-management-backend/cmd/table/domain"
	domainUser "restaurant-management-backend/cmd/user/domain"
)

type OrderResponse struct {
	Id        *types.OrderId        `json:"id"`
	Table     *domainTable.Table    `json:"table"`
	User      *domainUser.User      `json:"user"`
	Total     *types.OrderTotal     `json:"total"`
	CreatedAt *types.OrderCreatedAt `json:"created_at"`
	UpdatedAt *types.OrderUpdatedAt `json:"updated_at"`
}

func NewOrderResponse(id *types.OrderId, table *domainTable.Table, user *domainUser.User, total *types.OrderTotal, createdAt *types.OrderCreatedAt, updatedAt *types.OrderUpdatedAt) *OrderResponse {
	return &OrderResponse{
		Id:        id,
		Table:     table,
		User:      user,
		Total:     total,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
