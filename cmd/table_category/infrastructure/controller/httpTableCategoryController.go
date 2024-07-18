package controller

import (
	"encoding/json"
	"net/http"
	"restaurant-management-backend/cmd/shared/infrastructure"
	"restaurant-management-backend/cmd/table_category/domain"
	"strconv"

	"github.com/gorilla/mux"
)

type HttpTableCategoryController struct {
	serviceContainer *infrastructure.ServiceContainer
}

func NewHttpTableCategoryController(serviceContainer *infrastructure.ServiceContainer) *HttpTableCategoryController {
	return &HttpTableCategoryController{serviceContainer: serviceContainer}
}

func (this *HttpTableCategoryController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tableCategory domain.TableCategoryPrimitive
	if err := json.NewDecoder(r.Body).Decode(&tableCategory); err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := tableCategory.Validate(); err != nil {
		infrastructure.RespondValidationError(w, err)
		return
	}

	if err := this.serviceContainer.TableCategory.Create.Execute(tableCategory.Id, tableCategory.Name); err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, "Error al crear la categoría de tabla")
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, nil)
}

func (this *HttpTableCategoryController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tableCategories, err := this.serviceContainer.TableCategory.GetAll.Execute()
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	primitiveTableCategories := make([]*domain.TableCategoryPrimitive, len(tableCategories))
	for i, tableCategory := range tableCategories {
		primitiveTableCategories[i] = tableCategory.MapToPrimitive()
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, primitiveTableCategories)
}

func (this *HttpTableCategoryController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, "Invalid table category ID")
		return
	}

	tableCategory, err := this.serviceContainer.TableCategory.GetById.Execute(id)
	if err != nil {
		if _, ok := err.(*domain.TableCategoryNotFound); ok {
			infrastructure.RespondWithError(w, http.StatusNotFound, err.Error())
		} else {
			infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al obtener la categoría de tabla")
		}
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, tableCategory)
}

func (this *HttpTableCategoryController) Edit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, "Invalid table category ID")
		return
	}

	var tableCategory domain.TableCategoryPrimitive
	if err := json.NewDecoder(r.Body).Decode(&tableCategory); err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := tableCategory.Validate(); err != nil {
		infrastructure.RespondValidationError(w, err)
		return
	}

	if err := this.serviceContainer.TableCategory.Edit.Execute(id, tableCategory.Name); err != nil {
		if _, ok := err.(*domain.TableCategoryNotFound); ok {
			infrastructure.RespondWithError(w, http.StatusNotFound, err.Error())
		} else {
			infrastructure.RespondWithError(w, http.StatusBadRequest, "Error al editar la categoría de tabla")
		}
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, nil)
}

func (this *HttpTableCategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, "Invalid table category ID")
		return
	}

	if err := this.serviceContainer.TableCategory.Delete.Execute(id); err != nil {
		if _, ok := err.(*domain.TableCategoryNotFound); ok {
			infrastructure.RespondWithError(w, http.StatusNotFound, err.Error())
		} else {
			infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al eliminar la categoría de tabla")
		}
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, nil)
}
