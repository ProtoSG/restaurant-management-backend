package application

import (
	"restaurant-management-backend/cmd/table_category/domain"
	"restaurant-management-backend/cmd/table_category/domain/repository"
)

type TableCategoryGetAll struct {
	repository repository.TableCategoryRepository
}

func NewTableCategoryGetAll(repository repository.TableCategoryRepository) *TableCategoryGetAll {
	return &TableCategoryGetAll{
		repository: repository,
	}
}

func (this TableCategoryGetAll) Execute() ([]*domain.TableCategory, error) {
	return this.repository.GetAll()
}
