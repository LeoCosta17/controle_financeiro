package handlers

import (
	"app/internal/models"
	"app/internal/services"
	"encoding/json"
	"net/http"
)

type CostumersHandler struct {
	services services.Services
}

func (c *CostumersHandler) Create(w http.ResponseWriter, r *http.Request) {

	var costumer models.Costumer

	if err := json.NewDecoder(r.Body).Decode(&costumer); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	costumer, err := c.services.Costumers.Create(costumer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	data, err := json.Marshal(costumer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(data))
}

func (c *CostumersHandler) GetAll(w http.ResponseWriter, r *http.Request)
