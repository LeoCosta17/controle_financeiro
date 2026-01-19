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
	Costumers interface {
		Create(models.Costumer) (models.Costumer, error)
		GetAll() ([]models.Costumer, error)
	}
}

func NewServices(repository repositories.Repository) Services {
	return Services{
		Users:     &UserService{repository},
		Suppliers: &SupplierService{repository},
		Costumers: &CostumerServices{repository},
	}
}
