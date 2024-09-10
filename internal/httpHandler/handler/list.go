package handler

import (
	"encoding/json"
	"github.com/DeMarDeXis/VProj/internal/model"
	"net/http"
	"time"
)

func (h *Handler) createList(w http.ResponseWriter, r *http.Request) {
	userID, err := h.getUserId(r)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	var input model.TodoList
	var inputDate model.DateInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, err.Error())
		return
	}

	// Десериализация даты в структуру DateInput
	if err := json.Unmarshal(input.DoeDate, &inputDate); err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid date format")
		return
	}

	dueDate := time.Date(inputDate.Year, time.Month(inputDate.Month), inputDate.Day, 0, 0, 0, 0, time.UTC)
	input.DoeDate, _ = json.Marshal(dueDate.Format(time.RFC3339))

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

func (h *Handler) getAllLists(w http.ResponseWriter, r *http.Request) {
	userID, err := h.getUserId(r)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	lists, err := h.service.TodoList.GetAll(userID)

	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lists)
}
