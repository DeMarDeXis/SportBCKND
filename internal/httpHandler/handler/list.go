package handler

import (
	"encoding/json"
	"github.com/DeMarDeXis/VProj/internal/model"
	"net/http"
)

func (h Handler) createList(w http.ResponseWriter, r *http.Request) {
	userID, err := h.getUserId(r)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	var input model.TodoList

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.TodoList.Create(userID, input)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{"id": id}
	json.NewEncoder(w).Encode(response)
}
