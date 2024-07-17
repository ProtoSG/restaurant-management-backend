package repository

import (
	"restaurant-management-backend/cmd/table/domain"
	"restaurant-management-backend/cmd/table/domain/types"
)

type TableRepository interface {
	Create(table *domain.Table) error
	GetAll() ([]*domain.Table, error)
	GetById(id *types.TableId) (*domain.Table, error)
	Edit(table *domain.Table) error
	Delete(id *types.TableId) error
}
