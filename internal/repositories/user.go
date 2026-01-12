package repositories

import (
	"app/internal/models"
	"context"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func (r UserRepository) Create(ctx context.Context, user models.User) (models.User, error) {

	return models.User{}, nil
}

func (r UserRepository) GetAll() {

}

func (r UserRepository) GetByID() {

}

func (r UserRepository) Update() {

}

func (r UserRepository) Delete() {

}
