package repositories

import (
	"app/internal/models"
	"database/sql"
)

type ProductRepository struct {
	db *sql.DB
}

func (r *ProductRepository) Create(product models.Product) (models.Product, error) {

	query := `
		INSERT INTO products (name, description, storage_unit, unit_value)
		VALUES ($1,$2, $3, $4)
		RETURNING id
	`

	if err := r.db.QueryRow(
		query,
		product.Name,
		product.Description,
		product.StorageUnit,
		product.UnitValue,
	).Scan(&product.ID); err != nil {
		return models.Product{}, err
	}

	return product, nil
}
