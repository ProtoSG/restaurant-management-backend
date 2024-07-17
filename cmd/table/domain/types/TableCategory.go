package types

type TableCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewTableCategory(id int, name string) *TableCategory {
	return &TableCategory{
		ID:   id,
		Name: name,
	}
}
