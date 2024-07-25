package types

type OrderItemDescription struct {
	Value string `json:"value"`
}

func NewOrderItemDescription(value string) (*OrderItemDescription, error) {
	return &OrderItemDescription{Value: value}, nil
}
