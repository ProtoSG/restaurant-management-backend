package domain

import (
	"fmt"
	"restaurant-management-backend/cmd/table_category/domain/types"
)

type TableCategoryNotFound struct {
	TableCategoryId types.TableCategoryId
}

func NewTableCategoryNotFound(tableCategoryId types.TableCategoryId) *TableCategoryNotFound {
	return &TableCategoryNotFound{TableCategoryId: tableCategoryId}
}

func (this TableCategoryNotFound) Error() string {
	return fmt.Sprintf("TablaCategory con el ID %d no existe", this.TableCategoryId)
}
