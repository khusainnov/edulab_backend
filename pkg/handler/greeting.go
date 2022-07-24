package handler

import "net/http"

func (h *Handler) Greeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Greeting page"))
}
