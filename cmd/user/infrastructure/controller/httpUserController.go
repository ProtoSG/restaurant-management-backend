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
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := req.Validate(); err != nil {
		respondValidationError(w, err)
		return
	}

	if err := c.serviceContainer.User.Create.Execute(req.Id, req.Username, req.Password, req.Role); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithSuccess(w, http.StatusCreated, nil)
}

func (c *HttpUserController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := c.serviceContainer.User.GetAll.Execute()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	primitiveUsers := make([]*domain.UserPrimitive, len(users))
	for i, user := range users {
		primitiveUsers[i] = user.MapToPrimitive()
	}

	respondWithSuccess(w, http.StatusOK, primitiveUsers)
}

func (c *HttpUserController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	log.Println(id)

	user, erro := c.serviceContainer.User.GetById.Execute(id)
	if erro != nil {
		if _, ok := erro.(*domain.UserNotFoundError); ok {
			respondWithError(w, http.StatusNotFound, erro.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, erro.Error())
		}
		return
	}

	respondWithSuccess(w, http.StatusOK, user.MapToPrimitive())
}

func (c *HttpUserController) Edit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var req domain.UserPrimitive
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := req.Validate(); err != nil {
		respondValidationError(w, err)
		return
	}

	if err := c.serviceContainer.User.Edit.Execute(id, req.Username, req.Password, req.Role); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(w, http.StatusOK, nil)
}

func (c *HttpUserController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	if err := c.serviceContainer.User.Delete.Execute(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(w, http.StatusOK, nil)
}

// Función auxiliar para manejar las respuestas de éxito
func respondWithSuccess(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func respondValidationError(w http.ResponseWriter, validationError *domain.UserValidationError) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"field": validationError.Field, "message": validationError.Message})
}

// Función auxiliar para manejar las respuestas de error
func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
