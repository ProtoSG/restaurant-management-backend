package repository

import (
	"restaurant-management-backend/cmd/item_category/domain"
	"restaurant-management-backend/cmd/item_category/domain/types"
)

type ItemCategoryRepository interface {
	GetAll() ([]*domain.ItemCategory, error)
	GetById(id *types.ItemCategoryId) (*domain.ItemCategory, error)
	Create(itemCategory *domain.ItemCategory) error
	Edit(itemCategory *domain.ItemCategory) error
	Delete(id *types.ItemCategoryId) error
}
