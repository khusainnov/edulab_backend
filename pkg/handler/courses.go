package handler

import "net/http"

func (h *Handler) Courses(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Courses page"))
}
