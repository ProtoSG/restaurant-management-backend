package application

import (
	"restaurant-management-backend/cmd/table_category/domain"
	"restaurant-management-backend/cmd/table_category/domain/repository"
	"restaurant-management-backend/cmd/table_category/domain/types"
)

type TableCategoryEdit struct {
	repository repository.TableCategoryRepository
}

func NewTableCategoryEdit(repository repository.TableCategoryRepository) *TableCategoryEdit {
	return &TableCategoryEdit{
		repository: repository,
	}
}

func (this TableCategoryEdit) Execute(id int, name string) error {
	tableCategoryId, err := types.NewTableCategoryId(id)
	if err != nil {
		return err
	}

	tableCategoryName, err := types.NewTableCategoryName(name)
	if err != nil {
		return err
	}

	tableCategory := domain.NewTableCategory(tableCategoryId, tableCategoryName)
	return this.repository.Edit(tableCategory)
}
