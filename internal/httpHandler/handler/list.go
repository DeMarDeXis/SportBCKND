package handler

import (
	"encoding/json"
	"github.com/DeMarDeXis/VProj/internal/model"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func (h *Handler) createList(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) getListByID(w http.ResponseWriter, r *http.Request) {
	userID, err := h.getUserId(r)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	idList := chi.URLParam(r, "id")
	if idList == "" {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "ID is empty r invalid")
		return
	}

	id, err := strconv.Atoi(idList)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "ID is empty or invalid<Double>")
		return
	}

	list, err := h.service.TodoList.GetByID(userID, id)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(list)
}

func (h *Handler) updateList(w http.ResponseWriter, r *http.Request) {
	userID, err := h.getUserId(r)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	idList := chi.URLParam(r, "id")
	if idList == "" {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "ID is empty r invalid")
		return
	}

	id, err := strconv.Atoi(idList)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "ID is empty or invalid<Double>")
		return
	}

	var input model.UpdateListInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.TodoList.Update(userID, id, input); err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := statusResponse{Status: "ok"}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) deleteList(w http.ResponseWriter, r *http.Request) {
	userID, err := h.getUserId(r)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	idList := chi.URLParam(r, "id")
	if idList == "" {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "ID is empty r invalid")
		return
	}

	id, err := strconv.Atoi(idList)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "ID is empty or invalid<Double>")
		return
	}

	err = h.service.TodoList.Delete(userID, id)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := statusResponse{Status: "ok"}
	json.NewEncoder(w).Encode(response)
}
