package repositories

import (
	"app/internal/models"
	"database/sql"
)

type SupplierRepository struct {
	db *sql.DB
}

func (s *SupplierRepository) Create(supplier models.Supplier) (models.Supplier, error) {

	query := `
		INSERT INTO suppliers (name, email, telephone, document)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	if err := s.db.QueryRow(
		query,
		supplier.Name,
		supplier.Email,
		supplier.Telephone,
		supplier.Document,
	).Scan(
		&supplier.ID,
	); err != nil {
		return models.Supplier{}, err
	}

	return supplier, nil
}
