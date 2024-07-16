package types

type UserId struct {
	Value int `json:"value"`
}

func NewUserId(value int) (*UserId, error) {
	userId := UserId{Value: value}
	return &userId, nil
}

func (this UserId) ToValue() int {
	return this.Value
}
