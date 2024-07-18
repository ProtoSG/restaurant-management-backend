package domain

import (
	domainTable "restaurant-management-backend/cmd/table/domain"
	domainUser "restaurant-management-backend/cmd/user/domain"
	"time"
)

type OrderResponsePrimitive struct {
	Id        int                         `json:"id"`
	Table     *domainTable.TablePrimitive `json:"table"`
	User      *domainUser.UserPrimitive   `json:"user"`
	CreatedAt time.Time                   `json:"createdAt"`
	UpdatedAt time.Time                   `json:"updatedAt"`
}

func (this OrderResponse) MapToPrimitive() *OrderResponsePrimitive {
	return &OrderResponsePrimitive{
		Id:        this.Id.Value,
		Table:     this.Table.MapToPrimitive(),
		User:      this.User.MapToPrimitive(),
		CreatedAt: this.CreatedAt.Value,
		UpdatedAt: this.UpdatedAt.Value,
	}
}
