package domain

import (
	domainOrderItem "restaurant-management-backend/cmd/order_item/domain"
	domainTable "restaurant-management-backend/cmd/table/domain"
	domainUser "restaurant-management-backend/cmd/user/domain"
	"time"
)

type OrderResponsePrimitive struct {
	Id         int                                           `json:"id"`
	TableId    int                                           `json:"table_id"`
	Table      *domainTable.TableResponsePrimitive           `json:"table"`
	UserId     int                                           `json:"user_id"`
	User       *domainUser.UserPrimitive                     `json:"user"`
	OrderItems []*domainOrderItem.OrderItemResponsePrimitive `json:"order"`
	CreatedAt  time.Time                                     `json:"created_at"`
	UpdatedAt  time.Time                                     `json:"updated_at"`
}

func (this OrderResponse) MapToPrimitive() *OrderResponsePrimitive {
	return &OrderResponsePrimitive{
		Id:         this.Id.Value,
		TableId:    this.TableId.Value,
		Table:      this.Table.MapToPrimitive(),
		UserId:     this.UserId.Value,
		User:       this.User.MapToPrimitive(),
		OrderItems: mapOrderItemsToPrimitive(this.OrderItems),
		CreatedAt:  this.CreatedAt.Value,
		UpdatedAt:  this.UpdatedAt.Value,
	}
}

func mapOrderItemsToPrimitive(items []*domainOrderItem.OrderItemResponse) []*domainOrderItem.OrderItemResponsePrimitive {
	primitiveItems := make([]*domainOrderItem.OrderItemResponsePrimitive, len(items))
	for i, item := range items {
		primitiveItems[i] = item.MapToPrimitive()
	}
	return primitiveItems
}
