package types

type TableCategoryId struct {
	Value int `json:"value"`
}

func NewTableCategoryId(value int) (*TableCategoryId, error) {
	tableCategoryId := TableCategoryId{Value: value}
	return &tableCategoryId, nil
}

func (this TableCategoryId) ToValue() int {
	return this.Value
}
