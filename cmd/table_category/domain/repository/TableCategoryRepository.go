package repository

import (
	"restaurant-management-backend/cmd/table_category/domain"
	"restaurant-management-backend/cmd/table_category/domain/types"
)

type TableCategoryRepository interface {
	GetById(id *types.TableCategoryId) (*domain.TableCategory, error)
	GetAll() ([]*domain.TableCategory, error)
	Create(tableCategory *domain.TableCategory) error
	Edit(tableCategory *domain.TableCategory) error
	Delete(id *types.TableCategoryId) error
}
