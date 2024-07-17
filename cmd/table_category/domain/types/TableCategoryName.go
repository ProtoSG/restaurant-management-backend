package types

type TableCategoryName struct {
	Value string `json:"value"`
}

func NewTableCategoryName(value string) (*TableCategoryName, error) {
	tableCategoryName := TableCategoryName{Value: value}
	return &tableCategoryName, nil
}

func (this *TableCategoryName) ToValue() string {
	return this.Value
}
