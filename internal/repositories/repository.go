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
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		Users: &UserRepository{db},
	}
}
