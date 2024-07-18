package repository

import (
	"restaurant-management-backend/cmd/table/domain"
	"restaurant-management-backend/cmd/table/domain/types"
)

type TableRepository interface {
	Create(table *domain.Table) error
	GetAll() ([]*domain.TableResponse, error)
	GetById(id *types.TableId) (*domain.TableResponse, error)
	Edit(table *domain.Table) error
	Delete(id *types.TableId) error
}
