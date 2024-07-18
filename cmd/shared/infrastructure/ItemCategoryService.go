package infrastructure

import (
	"database/sql"
	"restaurant-management-backend/cmd/item_category/application"
	"restaurant-management-backend/cmd/item_category/infrastructure"
)

type ItemCategoryService struct {
	Create  *application.ItemCategoryCreate
	GetAll  *application.ItemCategoryGetAll
	GetById *application.ItemCategoryGetById
	Delete  *application.ItemCategoryDelete
	Edit    *application.ItemCategoryEdit
}

func NewItemCategoryService(db *sql.DB) ItemCategoryService {
	itemCategoryContainer := infrastructure.NewSQLiteItemCategoryRepository(db)
	return ItemCategoryService{
		Create:  application.NewItemCategoryCreate(itemCategoryContainer),
		GetAll:  application.NewItemCategoryGetAll(itemCategoryContainer),
		GetById: application.NewItemCategoryGetById(itemCategoryContainer),
		Delete:  application.NewItemCategoryDelete(itemCategoryContainer),
		Edit:    application.NewItemCategoryEdit(itemCategoryContainer),
	}
}
