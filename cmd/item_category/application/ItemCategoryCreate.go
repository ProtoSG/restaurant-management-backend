package application

import (
	"restaurant-management-backend/cmd/item_category/domain"
	"restaurant-management-backend/cmd/item_category/domain/repository"
	"restaurant-management-backend/cmd/item_category/domain/types"
)

type ItemCategoryCreate struct {
	repository repository.ItemCategoryRepository
}

func NewItemCategoryCreate(repository repository.ItemCategoryRepository) *ItemCategoryCreate {
	return &ItemCategoryCreate{repository: repository}
}

func (this *ItemCategoryCreate) Execute(id int, name string) error {
	itemCategoryId, err := types.NewItemCategoryId(id)
	if err != nil {
		return err
	}

	itemCategoryName, err := types.NewItemCategoryName(name)
	if err != nil {
		return err
	}

	itemCategory := domain.NewItemCategory(itemCategoryId, itemCategoryName)
	return this.repository.Create(itemCategory)
}
