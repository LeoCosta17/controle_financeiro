package services

import (
	"app/internal/models"
	"app/internal/repositories"
)

type SupplierService struct {
	repo repositories.Repository
}

func (s *SupplierService) Create(supplier models.Supplier) (models.Supplier, error) {

	if err := supplier.ValidateSupplier(); err != nil {
		return models.Supplier{}, err
	}

	supplier, err := s.repo.Suppliers.Create(supplier)
	if err != nil {
		return models.Supplier{}, err
	}

	return supplier, nil
}
