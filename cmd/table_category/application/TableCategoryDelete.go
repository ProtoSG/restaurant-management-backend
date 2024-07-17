package application

import (
	"restaurant-management-backend/cmd/table_category/domain/repository"
	"restaurant-management-backend/cmd/table_category/domain/types"
)

type TableCategoryDelete struct {
	repository repository.TableCategoryRepository
}

func NewTableCategoryDelete(repository repository.TableCategoryRepository) *TableCategoryDelete {
	return &TableCategoryDelete{
		repository: repository,
	}
}

func (this TableCategoryDelete) Execute(id int) error {
	tableCategoryId, err := types.NewTableCategoryId(id)
	if err != nil {
		return err
	}
	return this.repository.Delete(tableCategoryId)
}
