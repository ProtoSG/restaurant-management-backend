package application

import (
	"restaurant-management-backend/cmd/item_category/domain"
	"restaurant-management-backend/cmd/item_category/domain/repository"
	"restaurant-management-backend/cmd/item_category/domain/types"
)

type ItemCategoryGetById struct {
	repository repository.ItemCategoryRepository
}

func NewItemCategoryGetById(repository repository.ItemCategoryRepository) *ItemCategoryGetById {
	return &ItemCategoryGetById{
		repository: repository,
	}
}

func (this *ItemCategoryGetById) Execute(id int) (*domain.ItemCategory, error) {
	itemCategoryId, err := types.NewItemCategoryId(id)
	if err != nil {
		return nil, err
	}

	itemCategory, err := this.repository.GetById(itemCategoryId)
	if err != nil {
		return nil, domain.NewItemCategoryNotFound(itemCategoryId)
	}

	return itemCategory, nil
}
