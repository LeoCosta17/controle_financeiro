package services

import (
	"app/internal/models"
	"app/internal/repositories"
)

type Services struct {
	Users interface {
		Create(models.User) (models.User, error)
		GetAll() ([]models.User, error)
	}
	Suppliers interface {
		Create(models.Supplier) (models.Supplier, error)
	}
}

func NewServices(repository repositories.Repository) Services {
	return Services{
		Users:     &UserService{repository},
		Suppliers: &SupplierService{repository},
	}
}
