package application

import (
	"restaurant-management-backend/cmd/table_category/domain"
	"restaurant-management-backend/cmd/table_category/domain/repository"
	"restaurant-management-backend/cmd/table_category/domain/types"
)

type TableCategoryCreate struct {
	repository repository.TableCategoryRepository
}

func NewTableCategoryCreate(repository repository.TableCategoryRepository) *TableCategoryCreate {
	return &TableCategoryCreate{repository: repository}
}

func (this TableCategoryCreate) Execute(id int, name string) error {
	tableCategortyId, err := types.NewTableCategoryId(id)
	if err != nil {
		return err
	}

	tableCategoryName, err := types.NewTableCategoryName(name)
	if err != nil {
		return err
	}

	tableCategory := domain.NewTableCategory(tableCategortyId, tableCategoryName)
	return this.repository.Create(tableCategory)
}
