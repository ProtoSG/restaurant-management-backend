package infrastructure

import (
	"database/sql"
	"restaurant-management-backend/cmd/table/application"
	"restaurant-management-backend/cmd/table/infrastructure"
)

type TableService struct {
	Create  *application.TableCreate
	GetAll  *application.TableGetAll
	GetById *application.TableGetById
	Delete  *application.TableDelete
	Edit    *application.TableEdit
}

func NewTableService(db *sql.DB) TableService {
	tableContainer := infrastructure.NewSQLiteTableRepository(db)
	return TableService{
		Create:  application.NewTableCreate(tableContainer),
		GetAll:  application.NewTableGetAll(tableContainer),
		GetById: application.NewTableGetById(tableContainer),
		Delete:  application.NewTableDelete(tableContainer),
		Edit:    application.NewTableEdit(tableContainer),
	}
}
