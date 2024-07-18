package types

type ItemCategoryName struct {
	Value string `json:"value"`
}

func NewItemCategoryName(value string) (*ItemCategoryName, error) {
	itemCategoryName := ItemCategoryName{Value: value}
	return &itemCategoryName, nil
}

func (this ItemCategoryName) ToValue() string {
	return this.Value
}
