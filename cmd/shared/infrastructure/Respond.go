package infrastructure

import (
	"encoding/json"
	"net/http"
	"restaurant-management-backend/cmd/shared/domain"
)

func RespondWithSuccess(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func RespondValidationError(w http.ResponseWriter, validationError *domain.ValidationFieldError) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"field": validationError.Field, "message": validationError.Message})
}

func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
