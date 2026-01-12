package handlers

import (
	"app/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		log.Print(err)
		return
	}
}
