package handler

import (
	"github.com/gorilla/mux"
	"github.com/khusainnov/edulab/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", h.Greeting)
	r.HandleFunc("/courses", h.Courses)

	auth := r.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/sign-up", h.SignUp).Methods("POST")
		auth.HandleFunc("/sign-in", h.SignIn).Methods("GET", "POST")
	}

	api := r.PathPrefix("/api").Subrouter()
	api.Use(h.UserIdentity)
	{
		api.HandleFunc("/profile", h.Profile)

		profile := api.PathPrefix("/profile").Subrouter()
		//profile.Use()
		{
			profile.HandleFunc("/settings", h.Settings)
		}
	}

	return r
}
