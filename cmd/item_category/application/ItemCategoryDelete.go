package application

import (
	"restaurant-management-backend/cmd/item_category/domain"
	"restaurant-management-backend/cmd/item_category/domain/repository"
	"restaurant-management-backend/cmd/item_category/domain/types"
)

type ItemCategoryDelete struct {
	repository repository.ItemCategoryRepository
}

func NewItemCategoryDelete(repository repository.ItemCategoryRepository) *ItemCategoryDelete {
	return &ItemCategoryDelete{repository: repository}
}

func (this *ItemCategoryDelete) Execute(id int) error {
	itemCategoryId, err := types.NewItemCategoryId(id)
	if err != nil {
		return err
	}

	if itemCategory, _ := this.repository.GetById(itemCategoryId); itemCategory == nil {
		return domain.NewItemCategoryNotFound(*&itemCategoryId)
	}

	return this.repository.Delete(itemCategoryId)
}
