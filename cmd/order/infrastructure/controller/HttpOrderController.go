package controller

import (
	"encoding/json"
	"net/http"
	"restaurant-management-backend/cmd/order/domain"
	"restaurant-management-backend/cmd/shared/infrastructure"
	domainTable "restaurant-management-backend/cmd/table/domain"
	domainUser "restaurant-management-backend/cmd/user/domain"
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
		infrastructure.RespondWithError(w, http.StatusInternalServerError, "Error al obtener las ordenes")
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

		order.Table = table
		order.User = user

		primitiveResponseOrders[i] = order.MapToPrimitive()
	}

	infrastructure.RespondWithSuccess(w, http.StatusOK, primitiveResponseOrders)
}
