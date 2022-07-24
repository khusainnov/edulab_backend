package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) Profile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Profile\n"))
	id := r.Context().Value("userId")
	json.NewEncoder(w).Encode(&map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) Settings(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Settings"))
}
