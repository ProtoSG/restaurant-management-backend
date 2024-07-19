package domain

import (
	"restaurant-management-backend/cmd/order/domain/types"
	domainTable "restaurant-management-backend/cmd/table/domain"
	typesTable "restaurant-management-backend/cmd/table/domain/types"
	domainUser "restaurant-management-backend/cmd/user/domain"
	typesUser "restaurant-management-backend/cmd/user/domain/types"
)

type OrderResponse struct {
	Id        *types.OrderId             `json:"id"`
	TableId   *typesTable.TableId        `json:"table_id"`
	Table     *domainTable.TableResponse `json:"table"`
	UserId    *typesUser.UserId
	User      *domainUser.User      `json:"user"`
	Total     *types.OrderTotal     `json:"total"`
	CreatedAt *types.OrderCreatedAt `json:"created_at"`
	UpdatedAt *types.OrderUpdatedAt `json:"updated_at"`
}

func NewOrderResponse(
	id *types.OrderId,
	tableId *typesTable.TableId,
	table *domainTable.TableResponse,
	userId *typesUser.UserId,
	user *domainUser.User,
	total *types.OrderTotal,
	createdAt *types.OrderCreatedAt,
	updatedAt *types.OrderUpdatedAt,
) *OrderResponse {
	return &OrderResponse{
		Id:        id,
		TableId:   tableId,
		Table:     table,
		UserId:    userId,
		User:      user,
		Total:     total,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
