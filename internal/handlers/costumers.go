package handlers

import (
	httpresponse "app/internal/HTTPResponse"
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

	httpresponse.HTTPResponseOK(w, http.StatusCreated, costumer)
}

func (c *CostumersHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	costumers, err := c.services.Costumers.GetAll()
	if err != nil {
		httpresponse.HTTPResponseError(w, http.StatusInternalServerError, err)
		return
	}

	httpresponse.HTTPResponseOK(w, http.StatusOK, costumers)
}
