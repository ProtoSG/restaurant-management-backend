package controller

import (
	"encoding/json"
	"net/http"
	"restaurant-management-backend/cmd/item_category/domain"
	"restaurant-management-backend/cmd/shared/infrastructure"
	"strconv"

	"github.com/gorilla/mux"
)

type HttpItemCategoryController struct {
	serviceContainer *infrastructure.ServiceContainer
}

func NewHttpItemCategoryController(serviceContainer *infrastructure.ServiceContainer) *HttpItemCategoryController {
	return &HttpItemCategoryController{
		serviceContainer: serviceContainer,
	}
}

func (this *HttpItemCategoryController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var itemCategory domain.ItemCategoryPrimitive
	if err := json.NewDecoder(r.Body).Decode(&itemCategory); err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := itemCategory.Validate(); err != nil {
		infrastructure.RespondValidationError(w, err)
		return
	}

	if err := this.serviceContainer.ItemCategory.Create.Execute(itemCategory.Id, itemCategory.Name); err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, "Error al crear la categoría de producto")
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusCreated, nil)
}

func (this *HttpItemCategoryController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	itemCategories, err := this.serviceContainer.ItemCategory.GetAll.Execute()
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, "Error al obtener las categorías de producto")
		return
	}

	primitiveItemCategories := make([]*domain.ItemCategoryPrimitive, len(itemCategories))
	for i, itemCategory := range itemCategories {
		primitiveItemCategories[i] = itemCategory.MapToPrimitive()
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, primitiveItemCategories)
}

func (this *HttpItemCategoryController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, "Error al obtener la categoría de producto")
		return
	}

	itemCategory, err := this.serviceContainer.ItemCategory.GetById.Execute(id)
	if err != nil {
		if _, ok := err.(*domain.ItemCategoryNotFound); ok {
			infrastructure.RespondWithError(w, http.StatusNotFound, err.Error())
		} else {
			infrastructure.RespondWithError(w, http.StatusBadRequest, "Error al obtener la categoría de producto")
		}
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, itemCategory.MapToPrimitive())
}

func (this *HttpItemCategoryController) Edit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, "Error al obtener la categoría de producto")
		return
	}

	var itemCategory domain.ItemCategoryPrimitive
	if err := json.NewDecoder(r.Body).Decode(&itemCategory); err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := itemCategory.Validate(); err != nil {
		infrastructure.RespondValidationError(w, err)
		return
	}

	if err := this.serviceContainer.ItemCategory.Edit.Execute(id, itemCategory.Name); err != nil {
		if _, ok := err.(*domain.ItemCategoryNotFound); ok {
			infrastructure.RespondWithError(w, http.StatusNotFound, err.Error())
		} else {
			infrastructure.RespondWithError(w, http.StatusBadRequest, "Error al editar la categoría de producto")
		}
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, nil)
}

func (this *HttpItemCategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, "Error al obtener la categoría de producto")
		return
	}

	if err := this.serviceContainer.ItemCategory.Delete.Execute(id); err != nil {
		if _, ok := err.(*domain.ItemCategoryNotFound); ok {
			infrastructure.RespondWithError(w, http.StatusNotFound, err.Error())
		} else {
			infrastructure.RespondWithError(w, http.StatusBadRequest, "Error al eliminar la categoría de producto")
		}
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, nil)
}
