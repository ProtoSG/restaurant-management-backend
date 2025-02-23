package controller

import (
	"encoding/json"
	"net/http"
	domainInventory "restaurant-management-backend/cmd/inventory/domain"
	"restaurant-management-backend/cmd/order/domain"
	"restaurant-management-backend/cmd/shared/infrastructure"
	domainTable "restaurant-management-backend/cmd/table/domain"
	domainUser "restaurant-management-backend/cmd/user/domain"
	"strconv"

	"github.com/gorilla/mux"
)

type HttpOrderController struct {
	serviceContainer *infrastructure.ServiceContainer
}

func NewHttpOrderController(serviceContainer *infrastructure.ServiceContainer) *HttpOrderController {
	return &HttpOrderController{serviceContainer: serviceContainer}
}

func (this *HttpOrderController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var order domain.OrderPrimitive
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := order.Validate(); err != nil {
		infrastructure.RespondValidationError(w, err)
		return
	}

	if err := this.serviceContainer.Order.Create.Execute(
		order.Id,
		order.TableId,
		order.UserId,
		order.Total,
		order.CreatedAt,
		order.UpdatedAt,
		order.Completed,
	); err != nil {
		infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al crear la orden")
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, nil)
}

func (this *HttpOrderController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	orders, err := this.serviceContainer.Order.GetAll.Execute()
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	primitiveResponseOrders := make([]*domain.OrderResponsePrimitive, len(orders))
	for i, order := range orders {
		table, err := this.serviceContainer.Table.GetById.Execute(order.TableId.Value)
		if err != nil {
			if _, ok := err.(*domainTable.TableNotFound); ok {
				infrastructure.RespondWithError(w, http.StatusNotFound, err.Error())
			} else {
				infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al obtener la tabla")
			}
			return
		}

		user, err := this.serviceContainer.User.GetById.Execute(order.UserId.Value)
		if err != nil {
			if _, ok := err.(*domainUser.UserNotFoundError); ok {
				infrastructure.RespondWithError(w, http.StatusNotFound, err.Error())
			} else {
				infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al obtener la tabla")
			}
			return
		}

		order_items, err := this.serviceContainer.OrderItem.GetByOrder.Execute(order.Id.Value)
		if err != nil {
			infrastructure.RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		for _, order_item := range order_items {
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
		}

		order.Table = table
		order.User = user
		order.OrderItems = order_items

		primitiveResponseOrders[i] = order.MapToPrimitive()
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, primitiveResponseOrders)
}

func (this *HttpOrderController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	order, err := this.serviceContainer.Order.GetById.Execute(id)
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	table, err := this.serviceContainer.Table.GetById.Execute(order.TableId.Value)
	if err != nil {
		if _, ok := err.(*domainTable.TableNotFound); ok {
			infrastructure.RespondWithError(w, http.StatusNotFound, err.Error())
		} else {
			infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al obtener la tabla")
		}
		return
	}

	user, err := this.serviceContainer.User.GetById.Execute(order.UserId.Value)
	if err != nil {
		if _, ok := err.(*domainUser.UserNotFoundError); ok {
			infrastructure.RespondWithError(w, http.StatusNotFound, err.Error())
		} else {
			infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al obtener la tabla")
		}
		return
	}

	order_items, err := this.serviceContainer.OrderItem.GetByOrder.Execute(order.Id.Value)
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	for _, order_item := range order_items {
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
	}

	order.Table = table
	order.User = user
	order.OrderItems = order_items

	infrastructure.RespondWithSuccess(w, http.StatusOK, order.MapToPrimitive())
}

func (this *HttpOrderController) Edit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var order domain.OrderPrimitive
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := order.Validate(); err != nil {
		infrastructure.RespondValidationError(w, err)
		return
	}

	if err := this.serviceContainer.Order.Edit.Execute(
		id,
		order.TableId,
		order.UserId,
		order.Total,
		order.CreatedAt,
		order.UpdatedAt,
		order.Completed,
	); err != nil {
		infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al editar la orden")
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, nil)
}

func (this *HttpOrderController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		infrastructure.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := this.serviceContainer.Order.Delete.Execute(id); err != nil {
		infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al eliminar la orden")
		return
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, nil)

}
