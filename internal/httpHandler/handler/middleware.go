package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
)

func (h *Handler) userIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authHeader)
		if header == "" {
			newErrorResponse(w, h.logg, http.StatusUnauthorized, "empty auth header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			newErrorResponse(w, h.logg, http.StatusUnauthorized, "invalid auth header")
			return
		}

		userID, err := h.service.Auth.ParseToken(headerParts[1])
		if err != nil {
			newErrorResponse(w, h.logg, http.StatusUnauthorized, err.Error())
			return
		}

		r.Header.Set(userCtx, strconv.Itoa(userID))
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) getUserId(r *http.Request) (int, error) {
	id, ok := r.Context().Value(userCtx).(int)
	if !ok {
		return 0, errors.New("user id not found")
	}

	return id, nil
}
