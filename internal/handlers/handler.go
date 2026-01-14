package handlers

import (
	"app/internal/services"
	"net/http"
)

type Handler struct {
	Users interface {
		Create(http.ResponseWriter, *http.Request)
		GetAll(http.ResponseWriter, *http.Request)
	}
	Suppliers interface {
		Create(http.ResponseWriter, *http.Request)
	}
	Health interface {
		APIHealth(http.ResponseWriter, *http.Request)
	}
}

func NewHandlers(s services.Services) *Handler {
	return &Handler{
		Users:     &UserHandler{services: s},
		Suppliers: &SuppliersHandler{services: s},
		Health:    &HealtHanlder{},
	}
}
