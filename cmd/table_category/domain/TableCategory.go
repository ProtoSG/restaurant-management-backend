package domain

import "restaurant-management-backend/cmd/table_category/domain/types"

type TableCategory struct {
	Id   *types.TableCategoryId   `json:"id"`
	Name *types.TableCategoryName `json:"name"`
}

func NewTableCategory(id *types.TableCategoryId, name *types.TableCategoryName) *TableCategory {
	return &TableCategory{
		Id:   id,
		Name: name,
	}
}
