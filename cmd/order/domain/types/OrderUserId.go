package types

type OrderUserId struct {
	Value int `json:"value"`
}

func NewOrderUserId(value int) (*OrderUserId, error) {
	orderUserId := OrderUserId{Value: value}
	return &orderUserId, nil
}
