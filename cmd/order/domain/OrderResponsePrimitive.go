package domain

import (
	domainTable "restaurant-management-backend/cmd/table/domain"
	domainUser "restaurant-management-backend/cmd/user/domain"
	"time"
)

type OrderResponsePrimitive struct {
	Id        int                                 `json:"id"`
	TableId   int                                 `json:"table_id"`
	Table     *domainTable.TableResponsePrimitive `json:"table"`
	UserId    int                                 `json:"user_id"`
	User      *domainUser.UserPrimitive           `json:"user"`
	CreatedAt time.Time                           `json:"createdAt"`
	UpdatedAt time.Time                           `json:"updatedAt"`
}

func (this OrderResponse) MapToPrimitive() *OrderResponsePrimitive {
	return &OrderResponsePrimitive{
		Id:        this.Id.Value,
		TableId:   this.TableId.Value,
		Table:     this.Table.MapToPrimitive(),
		UserId:    this.UserId.Value,
		User:      this.User.MapToPrimitive(),
		CreatedAt: this.CreatedAt.Value,
		UpdatedAt: this.UpdatedAt.Value,
	}
}
