package controller

import (
	"encoding/json"
	"net/http"
	"restaurant-management-backend/cmd/shared/infrastructure"
	"restaurant-management-backend/cmd/table/domain"
	"strconv"

	"github.com/gorilla/mux"
)

type HttpTableController struct {
	serviceContainer *infrastructure.ServiceContainer
}

func NewHttpTableConstroller(serviceContainer *infrastructure.ServiceContainer) *HttpTableController {
	return &HttpTableController{serviceContainer: serviceContainer}
}

func (this *HttpTableController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var table domain.TablePrimitive
	if err := json.NewDecoder(r.Body).Decode(&table); err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := table.Validate(); err != nil {
		infrastructure.RespondValidationError(w, err)
		return
	}

	if err := this.serviceContainer.Table.Create.Execute(table.Id, table.Name, table.CategoryId, table.Status); err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, "Error al crear la tabla")
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, nil)
}

func (this *HttpTableController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tables, err := this.serviceContainer.Table.GetAll.Execute()
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al obtener las tablas")
		return
	}

	primitiveTables := make([]*domain.TableResponsePrimitive, len(tables))
	for i, table := range tables {
		primitiveTables[i] = table.MapToPrimitive()
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, primitiveTables)
}

func (this *HttpTableController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, "Invalid table ID")
		return
	}

	table, err := this.serviceContainer.Table.GetById.Execute(id)
	if err != nil {
		if _, ok := err.(*domain.TableNotFound); ok {
			infrastructure.RespondWithError(w, http.StatusNotFound, err.Error())
		} else {
			infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al obtener la tabla")
		}
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, table.MapToPrimitive())
}

func (this *HttpTableController) Edit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, "Invalid table ID")
		return
	}

	var table domain.TablePrimitive
	if err := json.NewDecoder(r.Body).Decode(&table); err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := table.Validate(); err != nil {
		infrastructure.RespondValidationError(w, err)
		return
	}

	if err := this.serviceContainer.Table.Edit.Execute(id, table.Name, table.CategoryId, table.Status); err != nil {
		if _, ok := err.(*domain.TableNotFound); ok {
			infrastructure.RespondWithError(w, http.StatusNotFound, err.Error())
		} else {
			infrastructure.RespondWithError(w, http.StatusBadRequest, err.Error())
		}
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, nil)
}

func (this *HttpTableController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, "Invalid table ID")
		return
	}

	if err := this.serviceContainer.Table.Delete.Execute(id); err != nil {
		if _, ok := err.(*domain.TableNotFound); ok {
			infrastructure.RespondWithError(w, http.StatusNotFound, err.Error())
		} else {
			infrastructure.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, nil)
}
