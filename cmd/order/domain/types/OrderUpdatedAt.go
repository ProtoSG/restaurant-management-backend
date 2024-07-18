package types

import "time"

type OrderUpdatedAt struct {
	Value time.Time `json:"value"`
}

func NewOrderUpdatedAt(value time.Time) (*OrderUpdatedAt, error) {
	orderUpdatedAt := OrderUpdatedAt{Value: value}
	return &orderUpdatedAt, nil
}
