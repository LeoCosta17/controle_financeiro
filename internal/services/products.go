package services

import (
	"app/internal/models"
	"app/internal/repositories"
)

type ProductServices struct {
	repo repositories.Repository
}

func (s *ProductServices) Create(product models.Product) (models.Product, error) {

	if err := product.ValidateProduct(); err != nil {
		return models.Product{}, err
	}

	product, err := s.repo.Products.Create(product)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}
