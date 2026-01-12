package repositories

import (
	"app/internal/models"
	"context"
	"database/sql"
)

type Repository struct {
	Users interface {
		Create(context.Context, models.User) (models.User, error)
	}
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		Users: &UserRepository{db},
	}
}
