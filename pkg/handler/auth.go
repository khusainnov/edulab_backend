package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/khusainnov/edulab/internal/entity/user"
	"github.com/sirupsen/logrus"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("signup\n"))
	var u user.User

	vars := mux.Vars(r)

	u = user.User{
		Name:     vars["name"],
		Surname:  vars["surname"],
		Username: vars["username"],
		Email:    vars["email"],
		Password: vars["password"],
		RoleName: vars["role_name"],
	}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		logrus.Errorf("%s", err.Error())
	}

	id, err := h.services.Authorization.CreateUser(u)
	if err != nil {
		logrus.Errorf("%s", err.Error())
	}

	err = json.NewEncoder(w).Encode(&map[string]interface{}{
		"id": id,
	})
	if err != nil {
		logrus.Errorf("Cannot encode id, %s", err.Error())
	}
}

type loginInfo struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("signin\n"))

	var u loginInfo

	vars := mux.Vars(r)

	u = loginInfo{
		Login:    vars["login"],
		Password: vars["password"],
	}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		logrus.Errorf("Cannot decode login info, due to error: %s", err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(u.Login, u.Password)
	if err != nil {
		logrus.Errorf("%s", err.Error())
		json.NewEncoder(w).Encode(&map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	err = json.NewEncoder(w).Encode(&map[string]interface{}{
		"token": token,
	})
	if err != nil {
		logrus.Errorf("Cannot encode token, due to error: %s", err.Error())
	}
}
