package types

type TableName struct {
	Value string `json:"value"`
}

func NewTableName(value string) (*TableName, error) {
	tableName := TableName{Value: value}
	return &tableName, nil
}

func (this TableName) ToValue() string {
	return this.Value
}
