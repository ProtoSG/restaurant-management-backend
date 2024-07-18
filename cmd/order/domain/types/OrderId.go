package types

type OrderId struct {
	Value int `json:"value"`
}

func NewOrderId(value int) (*OrderId, error) {
	orderId := OrderId{Value: value}
	return &orderId, nil
}
