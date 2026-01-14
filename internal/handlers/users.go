package handlers

import (
	"app/internal/models"
	"app/internal/services"
	"encoding/json"
	"log"
	"net/http"
)

type UserHandler struct {
	services services.Services
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Print(err)
		return
	}

	user, err := h.services.Users.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	data, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(data))

}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	users, err := h.services.Users.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	data, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}
