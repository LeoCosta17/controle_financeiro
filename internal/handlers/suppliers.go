package handlers

import (
	"app/internal/models"
	"app/internal/services"
	"encoding/json"
	"net/http"
)

type SuppliersHandler struct {
	services services.Services
}

func (s *SuppliersHandler) Create(w http.ResponseWriter, r *http.Request) {

	var supplier models.Supplier

	if err := json.NewDecoder(r.Body).Decode(&supplier); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	supplier, err := s.services.Suppliers.Create(supplier)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	data, err := json.Marshal(supplier)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(data))
}
