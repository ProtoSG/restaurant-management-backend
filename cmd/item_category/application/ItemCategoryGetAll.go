package application

import (
	"restaurant-management-backend/cmd/item_category/domain"
	"restaurant-management-backend/cmd/item_category/domain/repository"
)

type ItemCategoryGetAll struct {
	repository repository.ItemCategoryRepository
}

func NewItemCategoryGetAll(repository repository.ItemCategoryRepository) *ItemCategoryGetAll {
	return &ItemCategoryGetAll{repository: repository}
}

func (this *ItemCategoryGetAll) Execute() ([]*domain.ItemCategory, error) {
	return this.repository.GetAll()
}
