package application

import (
	"restaurant-management-backend/cmd/item_category/domain"
	"restaurant-management-backend/cmd/item_category/domain/repository"
	"restaurant-management-backend/cmd/item_category/domain/types"
)

type ItemCategoryEdit struct {
	repository repository.ItemCategoryRepository
}

func NewItemCategoryEdit(repository repository.ItemCategoryRepository) *ItemCategoryEdit {
	return &ItemCategoryEdit{repository: repository}
}

func (this *ItemCategoryEdit) Execute(id int, name string) error {
	itemCategoryId, err := types.NewItemCategoryId(id)
	if err != nil {
		return err
	}

	if itemCategory, _ := this.repository.GetById(itemCategoryId); itemCategory == nil {
		return domain.NewItemCategoryNotFound(*&itemCategoryId)
	}

	itemCategoryName, err := types.NewItemCategoryName(name)
	if err != nil {
		return err
	}

	itemCategory := domain.NewItemCategory(itemCategoryId, itemCategoryName)
	return this.repository.Edit(itemCategory)
}
