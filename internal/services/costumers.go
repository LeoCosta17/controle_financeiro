package services

import (
	"app/internal/models"
	"app/internal/repositories"
)

type CostumerServices struct {
	repo repositories.Repository
}

func (s *CostumerServices) Create(costumer models.Costumer) (models.Costumer, error) {

	if err := costumer.ValidateCostumer(); err != nil {
		return models.Costumer{}, err
	}

	costumer, err := s.repo.Costumers.Create(costumer)
	if err != nil {
		return models.Costumer{}, err
	}

	return costumer, nil
}
