package handlers

import (
	httpresponse "app/internal/HTTPResponse"
	"app/internal/models"
	"app/internal/services"
	"encoding/json"
	"net/http"
)

type ProductsHandler struct {
	services services.Services
}

func (h *ProductsHandler) Create(w http.ResponseWriter, r *http.Request) {

	var product models.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		httpresponse.HTTPResponseError(w, http.StatusBadRequest, err)
		return
	}

	product, err := h.services.Products.Create(product)
	if err != nil {
		httpresponse.HTTPResponseError(w, http.StatusInternalServerError, err)
	}

	httpresponse.HTTPResponseOK(w, http.StatusCreated, product)
}
