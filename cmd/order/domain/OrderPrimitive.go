package domain

import (
	"restaurant-management-backend/cmd/shared/domain"
	"time"
)

type OrderPrimitive struct {
	Id        int       `json:"id"`
	TableId   int       `json:"table_id"`
	UserId    int       `json:"user_id"`
	Total     float32   `json:"total"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (this Order) MapToPrimitive() *OrderPrimitive {
	return &OrderPrimitive{
		Id:        this.Id.Value,
		TableId:   this.TableId.Value,
		UserId:    this.UserId.Value,
		Total:     this.Total.Value,
		CreatedAt: this.CreatedAt.Value,
		UpdatedAt: this.UpdatedAt.Value,
	}
}

func (this OrderPrimitive) Validate() *domain.ValidationFieldError {
	if this.TableId == 0 {
		return &domain.ValidationFieldError{Field: "table_id", Message: "Table id is required"}
	}
	if this.UserId == 0 {
		return &domain.ValidationFieldError{Field: "user_id", Message: "User id is required"}
	}
	if this.Total == 0 {
		return &domain.ValidationFieldError{Field: "total", Message: "Total is required"}
	}
	if this.CreatedAt.IsZero() {
		return &domain.ValidationFieldError{Field: "created_at", Message: "Created at is required"}
	}
	if this.UpdatedAt.IsZero() {
		return &domain.ValidationFieldError{Field: "updated_at", Message: "Updated at is required"}
	}
	return nil
}
