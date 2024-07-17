package types

type TableId struct {
	Value int `json:"value"`
}

func NewTableId(value int) (*TableId, error) {
	tableId := TableId{Value: value}
	return &tableId, nil
}

func (this TableId) ToValue() int {
	return this.Value
}
