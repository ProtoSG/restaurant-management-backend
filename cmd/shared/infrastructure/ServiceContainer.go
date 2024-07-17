package infrastructure

import (
	"fmt"
	"log"
	"os"
	"restaurant-management-backend/cmd/user/application"
	repository "restaurant-management-backend/cmd/user/infrastructure"

	"github.com/joho/godotenv"
)

type ServiceContainer struct {
	User struct {
		Create  *application.UserCreate
		GetAll  *application.UserGetAll
		GetById *application.UserGetById
		Delete  *application.UserDelete
		Edit    *application.UserEdit
	}
}

func NewServiceContainer() *ServiceContainer {
	// NOTE: Elegimos usar la infraestructura que se desea
	// Cargar variables de entorno desde el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Obtener la URL de conexión a la base de datos desde las variables de entorno
	databaseURL := os.Getenv("TURSO_DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("TURSO_DATABASE_URL not found in .env file")
	}

	// Obtener el token de autenticación desde las variables de entorno
	authToken := os.Getenv("TURSO_AUTH_TOKEN")
	if authToken == "" {
		log.Fatal("TURSO_AUTH_TOKEN not found in .env file")
	}

	// Construir la URL de conexión a la base de datos
	url := fmt.Sprintf("%s?authToken=%s", databaseURL, authToken)
	println(url)
	userContainer := repository.NewSQLiteUserRepository(url)

	container := &ServiceContainer{}

	container.User.Create = application.NewUserCreate(userContainer)
	container.User.GetAll = application.NewUserGetAll(userContainer)
	container.User.GetById = application.NewUserGetById(userContainer)
	container.User.Delete = application.NewUserDelete(userContainer)
	container.User.Edit = application.NewUserEdit(userContainer)

	return container
}
