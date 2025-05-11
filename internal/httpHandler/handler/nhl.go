package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func (h *Handler) getTeams(w http.ResponseWriter, r *http.Request) {
	teams, err := h.service.NHLList.GetTeams()
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(teams)
}

func (h *Handler) getSchedule(w http.ResponseWriter, r *http.Request) {
	schedule, err := h.service.NHLList.GetSchedule()
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(schedule)
}

func (h *Handler) getLastSchedule(w http.ResponseWriter, r *http.Request) {
	count, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, err.Error())
		return
	}

	schedule, err := h.service.NHLList.GetLastSchedule(count)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(schedule)
}
