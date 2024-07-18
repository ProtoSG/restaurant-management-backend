package infrastructure

import (
	"database/sql"
	"restaurant-management-backend/cmd/user/application"
	"restaurant-management-backend/cmd/user/infrastructure"
)

type UserService struct {
	Create  *application.UserCreate
	GetAll  *application.UserGetAll
	GetById *application.UserGetById
	Delete  *application.UserDelete
	Edit    *application.UserEdit
}

func NewUserService(db *sql.DB) UserService {
	userContainer := infrastructure.NewSQLiteUserRepository(db)
	return UserService{
		Create:  application.NewUserCreate(userContainer),
		GetAll:  application.NewUserGetAll(userContainer),
		GetById: application.NewUserGetById(userContainer),
		Delete:  application.NewUserDelete(userContainer),
		Edit:    application.NewUserEdit(userContainer),
	}
}
