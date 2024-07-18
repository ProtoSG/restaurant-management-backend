package types

import "time"

type OrderCreatedAt struct {
	Value time.Time `json:"value"`
}

func NewOrderCreatedAt(value time.Time) (*OrderCreatedAt, error) {
	orderCreatedAt := OrderCreatedAt{Value: value}
	return &orderCreatedAt, nil
}
