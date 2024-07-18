package controller

import (
	"encoding/json"
	"net/http"
	"restaurant-management-backend/cmd/inventory/domain"
	"restaurant-management-backend/cmd/shared/infrastructure"
	"strconv"

	"github.com/gorilla/mux"
)

type HttpInventoryController struct {
	serviceContainer *infrastructure.ServiceContainer
}

func NewHttpInventoryController(serviceContainer *infrastructure.ServiceContainer) *HttpInventoryController {
	return &HttpInventoryController{serviceContainer: serviceContainer}
}

func (this *HttpInventoryController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var inventory domain.InventoryPrimitive
	if err := json.NewDecoder(r.Body).Decode(&inventory); err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := inventory.Validate(); err != nil {
		infrastructure.RespondValidationError(w, err)
		return
	}

	if err := this.serviceContainer.Inventory.Create.Execute(inventory.Id, inventory.Name, inventory.ItemCategoryId, inventory.Quantity, inventory.Price); err != nil {
		infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al crear el inventario")
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusCreated, nil)
}

func (this *HttpInventoryController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	inventory, err := this.serviceContainer.Inventory.GetAll.Execute()
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al obtener el inventario")
		return
	}

	primitiveInventory := make([]*domain.InventoryResponsePrimitive, len(inventory))
	for i, inv := range inventory {
		primitiveInventory[i] = inv.MapToPrimitive()
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, primitiveInventory)
}

func (this *HttpInventoryController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, "Falta el id")
		return
	}

	inventory, err := this.serviceContainer.Inventory.GetById.Execute(id)
	if err != nil {
		if _, ok := err.(*domain.InventoryNotFound); ok {
			infrastructure.RespondWithError(w, http.StatusNotFound, err.Error())
		} else {
			infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al obtener el inventario")
		}
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, inventory.MapToPrimitive())
}

func (this *HttpInventoryController) Edit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	var inventory domain.InventoryPrimitive
	if err := json.NewDecoder(r.Body).Decode(&inventory); err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := inventory.Validate(); err != nil {
		infrastructure.RespondValidationError(w, err)
		return
	}

	if err != this.serviceContainer.Inventory.Edit.Execute(id, inventory.Name, inventory.ItemCategoryId, inventory.Quantity, inventory.Price) {
		if _, ok := err.(*domain.InventoryNotFound); ok {
			infrastructure.RespondWithError(w, http.StatusNotFound, err.Error())
		} else {
			infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al editar el inventario")
		}
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, nil)
}

func (this *HttpInventoryController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, "Falta el id")
		return
	}

	err = this.serviceContainer.Inventory.Delete.Execute(id)
	if err != nil {
		if _, ok := err.(*domain.InventoryNotFound); ok {
			infrastructure.RespondWithError(w, http.StatusNotFound, err.Error())
		} else {
			infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al eliminar el inventario")
		}
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, nil)
}
