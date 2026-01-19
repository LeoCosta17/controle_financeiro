package repositories

import (
	"app/internal/models"
	"database/sql"
)

type Repository struct {
	Users interface {
		Create(models.User) (models.User, error)
		GetAll() ([]models.User, error)
	}
	Suppliers interface {
		Create(models.Supplier) (models.Supplier, error)
	}
	Costumers interface {
		Create(models.Costumer) (models.Costumer, error)
	}
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		Users:     &UserRepository{db: db},
		Suppliers: &SupplierRepository{db: db},
		Costumers: &CostumerRepository{db: db},
	}
}
