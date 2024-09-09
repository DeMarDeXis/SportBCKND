package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(w http.ResponseWriter, logg *slog.Logger, statusCode int, message string) {
	logg.Error(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := errorResponse{Message: message}
	json.NewEncoder(w).Encode(response)
}
