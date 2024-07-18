package types

type ItemCategoryId struct {
	Value int `json:"value"`
}

func NewItemCategoryId(value int) (*ItemCategoryId, error) {
	ItemCategoryId := ItemCategoryId{Value: value}
	return &ItemCategoryId, nil
}

func (this ItemCategoryId) ToValue() int {
	return this.Value
}
