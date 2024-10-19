package handler

import (
	"encoding/json"
	"github.com/DeMarDeXis/VProj/internal/model"
	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var input model.User

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, err.Error())
		return
	}

	if err := input.Validate(); err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Auth.CreateUser(input)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{"id": id}
	json.NewEncoder(w).Encode(response)
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var input signInInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.service.Auth.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{"token": token}
	json.NewEncoder(w).Encode(response)
}
