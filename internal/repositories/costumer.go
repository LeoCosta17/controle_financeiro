package repositories

import (
	"app/internal/models"
	"database/sql"
)

type CostumerRepository struct {
	db *sql.DB
}

func (r *CostumerRepository) Create(costumer models.Costumer) (models.Costumer, error) {

	query := `
		INSERT INTO costumers (name, email, telephone, document)
		VALUES ($1,$2,$3,$4)
		RETURNING id
	`

	if err := r.db.QueryRow(
		query,
		costumer.Name,
		costumer.Email,
		costumer.Telephone,
		costumer.Document,
	).Scan(
		&costumer.ID,
	); err != nil {
		return models.Costumer{}, err
	}

	return costumer, nil
}

func (r *CostumerRepository) GetAll() ([]models.Costumer, error) {
	query := `
		SELECT id, name, email, telephone, document FROM costumers
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var costumers []models.Costumer

	for rows.Next() {

		var costumer models.Costumer

		if err := rows.Scan(
			&costumer.ID,
			&costumer.Name,
			&costumer.Email,
			&costumer.Telephone,
			&costumer.Document,
		); err != nil {
			return nil, err
		}

		costumers = append(costumers, costumer)
	}

	return costumers, nil
}
