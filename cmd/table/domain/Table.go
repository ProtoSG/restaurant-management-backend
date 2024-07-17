package domain

import "restaurant-management-backend/cmd/table/domain/types"

type Table struct {
	Id        *types.TableId         `json:"id"`
	Name      *types.TableName       `json:"name"`
	CategoyId *types.TableCategoryId `json:"category_id"`
	Status    *types.TableStatus     `json:"status"`
}

func NewTable(id *types.TableId, name *types.TableName, category_id *types.TableCategoryId, status *types.TableStatus) *Table {
	return &Table{
		Id:        id,
		Name:      name,
		CategoyId: category_id,
		Status:    status,
	}
}

type TableResponse struct {
	Id      *types.TableId       `json:"id"`
	Name    *types.TableName     `json:"name"`
	Categoy *types.TableCategory `json:"category"`
	Status  *types.TableStatus   `json:"status"`
}

func NewTableResponse(id *types.TableId, name *types.TableName, category *types.TableCategory, status *types.TableStatus) *TableResponse {
	return &TableResponse{
		Id:      id,
		Name:    name,
		Categoy: category,
		Status:  status,
	}
}
