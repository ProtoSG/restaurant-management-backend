package application

import (
	"restaurant-management-backend/cmd/table_category/domain"
	"restaurant-management-backend/cmd/table_category/domain/repository"
	"restaurant-management-backend/cmd/table_category/domain/types"
)

type TableCategoryGetById struct {
	repository repository.TableCategoryRepository
}

func NewTableCategoryGetById(repository repository.TableCategoryRepository) *TableCategoryGetById {
	return &TableCategoryGetById{
		repository: repository,
	}
}

func (this TableCategoryGetById) Execute(id int) (*domain.TableCategory, error) {
	tableCategoryId, err := types.NewTableCategoryId(id)
	if err != nil {
		return nil, err
	}

	tableCategory, err := this.repository.GetById(tableCategoryId)
	if err != nil {
		return nil, domain.NewTableCategoryNotFound(*tableCategoryId)
	}

	return tableCategory, nil
}
