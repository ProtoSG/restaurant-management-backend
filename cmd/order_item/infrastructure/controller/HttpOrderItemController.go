package controller

import (
	"encoding/json"
	"net/http"
	domainInventory "restaurant-management-backend/cmd/inventory/domain"
	"restaurant-management-backend/cmd/order_item/domain"
	"restaurant-management-backend/cmd/shared/infrastructure"
	"strconv"

	"github.com/gorilla/mux"
)

type HttpOrderItemController struct {
	serviceContainer *infrastructure.ServiceContainer
}

func NewHttpOrderItemController(serviceContainer *infrastructure.ServiceContainer) *HttpOrderItemController {
	return &HttpOrderItemController{serviceContainer: serviceContainer}
}

func (this HttpOrderItemController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var order_item domain.OrderItemPrimitive
	if err := json.NewDecoder(r.Body).Decode(&order_item); err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := order_item.Validate(); err != nil {
		infrastructure.RespondValidationError(w, err)
		return
	}

	if err := this.serviceContainer.OrderItem.Create.Execute(
		order_item.Id,
		order_item.OrderId,
		order_item.ItemId,
		order_item.Quantity,
		order_item.SubTotal,
		order_item.Description,
		order_item.Takeaway,
	); err != nil {
		infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al crear el item de la orden")
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, nil)
}

func (this HttpOrderItemController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	orders_items, err := this.serviceContainer.OrderItem.GetAll.Execute()
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al obtener los items de las ordenes")
		return
	}

	primitiveResponseOrdersItems := make([]*domain.OrderItemResponsePrimitive, len(orders_items))
	for i, order_item := range orders_items {
		item, err := this.serviceContainer.Inventory.GetById.Execute(order_item.ItemId.Value)
		if err != nil {
			if _, ok := err.(*domainInventory.InventoryNotFound); ok {
				infrastructure.RespondWithError(w, http.StatusNotFound, err.Error())
			} else {
				infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al obtener el Item")
			}
			return
		}

		order_item.Item = item

		primitiveResponseOrdersItems[i] = order_item.MapToPrimitive()
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, primitiveResponseOrdersItems)
}

func (this HttpOrderItemController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	order_item, err := this.serviceContainer.OrderItem.GetById.Execute(id)
	if err != nil {
		if _, ok := err.(*domain.OrderNotFound); ok {
			infrastructure.RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error en el servidor")
		return
	}

	item, err := this.serviceContainer.Inventory.GetById.Execute(order_item.ItemId.Value)
	if err != nil {
		if _, ok := err.(*domainInventory.InventoryNotFound); ok {
			infrastructure.RespondWithError(w, http.StatusNotFound, err.Error())
		} else {
			infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al obtener el Item")
		}
		return
	}

	order_item.Item = item
	infrastructure.RespondWithSuccess(w, http.StatusOK, order_item.MapToPrimitive())
}

func (this HttpOrderItemController) Edit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	var order_item domain.OrderItemPrimitive
	if err := json.NewDecoder(r.Body).Decode(&order_item); err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := order_item.Validate(); err != nil {
		infrastructure.RespondValidationError(w, err)
		return
	}

	if err := this.serviceContainer.OrderItem.Edit.Execute(
		id,
		order_item.OrderId,
		order_item.ItemId,
		order_item.Quantity,
		order_item.SubTotal,
		order_item.Description,
		order_item.Takeaway,
	); err != nil {
		if _, ok := err.(*domain.OrderNotFound); ok {
			infrastructure.RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error en el servidor")
		return
	}
	infrastructure.RespondWithSuccess(w, http.StatusOK, nil)
}

func (this HttpOrderItemController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, "Id inválid")
		return
	}

	if err := this.serviceContainer.OrderItem.Delete.Execute(id); err != nil {
		if _, ok := err.(*domain.OrderNotFound); ok {
			infrastructure.RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error en el servidor")
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, nil)
}
