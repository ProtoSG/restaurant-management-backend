package infrastructure

import (
	tableApplication "restaurant-management-backend/cmd/table/application"
	tableRepository "restaurant-management-backend/cmd/table/infrastructure"
	userApplication "restaurant-management-backend/cmd/user/application"
	userRepository "restaurant-management-backend/cmd/user/infrastructure"
)

type ServiceContainer struct {
	User struct {
		Create  *userApplication.UserCreate
		GetAll  *userApplication.UserGetAll
		GetById *userApplication.UserGetById
		Delete  *userApplication.UserDelete
		Edit    *userApplication.UserEdit
	}

	Table struct {
		Create  *tableApplication.TableCreate
		GetAll  *tableApplication.TableGetAll
		GetById *tableApplication.TableGetById
		Delete  *tableApplication.TableDelete
		Edit    *tableApplication.TableEdit
	}
}

func NewServiceContainer() *ServiceContainer {

	env := NewEnv()
	userContainer := userRepository.NewSQLiteUserRepository(env.URL)

	container := &ServiceContainer{}

	container.User.Create = userApplication.NewUserCreate(userContainer)
	container.User.GetAll = userApplication.NewUserGetAll(userContainer)
	container.User.GetById = userApplication.NewUserGetById(userContainer)
	container.User.Delete = userApplication.NewUserDelete(userContainer)
	container.User.Edit = userApplication.NewUserEdit(userContainer)

	tableContainer := tableRepository.NewSQLiteTableRepository(env.URL)

	container.Table.Create = tableApplication.NewTableCreate(tableContainer)
	container.Table.GetAll = tableApplication.NewTableGetAll(tableContainer)
	container.Table.GetById = tableApplication.NewTableGetById(tableContainer)
	container.Table.Delete = tableApplication.NewTableDelete(tableContainer)
	container.Table.Edit = tableApplication.NewTableEdit(tableContainer)

	return container
}
