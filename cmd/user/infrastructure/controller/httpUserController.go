package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"restaurant-management-backend/cmd/shared/infrastructure"
	"restaurant-management-backend/cmd/user/domain"
	"strconv"

	"github.com/gorilla/mux"
)

type HttpUserController struct {
	serviceContainer *infrastructure.ServiceContainer
}

type req struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// Inicializar el Controlador
func NewHttpUserController(serviceContainer *infrastructure.ServiceContainer) *HttpUserController {
	return &HttpUserController{serviceContainer: serviceContainer}
}

// Métodos

func (c *HttpUserController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req domain.UserPrimitive
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := req.Validate(); err != nil {
		infrastructure.RespondValidationError(w, err)
		return
	}

	if err := c.serviceContainer.User.Create.Execute(req.Id, req.Username, req.Password, req.Role); err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, "Error al crear el usuario")
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusCreated, nil)
}

func (c *HttpUserController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := c.serviceContainer.User.GetAll.Execute()
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al obtener todos los usuarios")
		return
	}

	primitiveUsers := make([]*domain.UserPrimitive, len(users))
	for i, user := range users {
		primitiveUsers[i] = user.MapToPrimitive()
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, primitiveUsers)
}

func (c *HttpUserController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	log.Println(id)

	user, erro := c.serviceContainer.User.GetById.Execute(id)
	if erro != nil {
		if _, ok := erro.(*domain.UserNotFoundError); ok {
			infrastructure.RespondWithError(w, http.StatusNotFound, erro.Error())
		} else {
			infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al obtener el usuario")
		}
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, user.MapToPrimitive())
}

func (c *HttpUserController) Edit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var req domain.UserPrimitive
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := req.Validate(); err != nil {
		infrastructure.RespondValidationError(w, err)
		return
	}

	if err := c.serviceContainer.User.Edit.Execute(id, req.Username, req.Password, req.Role); err != nil {
		infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al editar el usuario")
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, nil)
}

func (c *HttpUserController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	if err := c.serviceContainer.User.Delete.Execute(id); err != nil {
		infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al eliminar el usuario")
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, nil)
}
