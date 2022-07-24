package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) UserIdentity(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)
		if header == "" {
			json.NewEncoder(w).Encode(&map[string]interface{}{
				"status":  http.StatusUnauthorized,
				"message": "empty authorization header",
			})
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			json.NewEncoder(w).Encode(&map[string]interface{}{
				"status":  http.StatusUnauthorized,
				"message": "invalid authorization header",
			})
			return
		}

		userId, err := h.services.Authorization.ParseToken(headerParts[1])
		if err != nil {
			json.NewEncoder(w).Encode(&map[string]interface{}{
				"status":  http.StatusUnauthorized,
				"message": err.Error(),
			})
			return
		}

		ctx := context.WithValue(r.Context(), "userId", userId)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}
