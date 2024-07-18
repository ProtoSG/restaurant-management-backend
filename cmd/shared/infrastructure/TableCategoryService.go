package infrastructure

import (
	"database/sql"
	"restaurant-management-backend/cmd/table_category/application"
	"restaurant-management-backend/cmd/table_category/infrastructure"
)

type TableCategoryService struct {
	Create  *application.TableCategoryCreate
	GetAll  *application.TableCategoryGetAll
	GetById *application.TableCategoryGetById
	Delete  *application.TableCategoryDelete
	Edit    *application.TableCategoryEdit
}

func NewTableCategoryService(db *sql.DB) TableCategoryService {
	tableCategoryContainer := infrastructure.NewSQLiteTableCategoryRepository(db)
	return TableCategoryService{
		Create:  application.NewTableCategoryCreate(tableCategoryContainer),
		GetAll:  application.NewTableCategoryGetAll(tableCategoryContainer),
		GetById: application.NewTableCategoryGetById(tableCategoryContainer),
		Delete:  application.NewTableCategoryDelete(tableCategoryContainer),
		Edit:    application.NewTableCategoryEdit(tableCategoryContainer),
	}
}
