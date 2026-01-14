package repositories

import (
	"app/internal/models"
	"database/sql"
	"errors"
)

type UserRepository struct {
	db *sql.DB
}

// Repositorio - Criar um novo usuário no banco de dados
func (r UserRepository) Create(user models.User) (models.User, error) {

	query := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	if err := r.db.QueryRow(
		query,
		user.Name,
		user.Email,
		user.Password,
	).Scan(
		&user.ID,
	); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r UserRepository) GetAll() ([]models.User, error) {
	query := `
		SELECT id, name, email FROM users
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
		); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, errors.New("no users found")
			}

			return nil, err
		}

		users = append(users, user)

	}

	return users, nil
}

func (r UserRepository) GetByID() {

}

func (r UserRepository) Update() {

}

func (r UserRepository) Delete() {

}
