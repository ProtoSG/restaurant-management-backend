package types

type OrderItemId struct {
	Value int `json:"value"`
}

func NewOrderItemId(value int) (*OrderItemId, error) {
	orderItemId := OrderItemId{Value: value}
	return &orderItemId, nil
}
