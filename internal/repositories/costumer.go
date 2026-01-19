package repositories

import (
	"app/internal/models"
	"database/sql"
)

type CostumerRepository struct {
	db *sql.DB
}

func (c *CostumerRepository) Create(costumer models.Costumer) (models.Costumer, error) {

	query := `
		INSERT INTO costumers (name, email, telephone, document)
		VALUES ($1,$2,$3,$4)
		RETURNING id
	`

	if err := c.db.QueryRow(
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
