package models

import (
	"app/internal/enums"
	"errors"
)

type Product struct {
	ID          uint64
	Name        string
	Description string
	StorageUnit enums.StorageUnit
	UnitValue   float64
}

func (p *Product) ValidateProduct() error {

	if err := validateProductData(p); err != nil {
		return err
	}

	return nil
}

func validateProductData(p *Product) error {
	if p.Name == "" {
		return errors.New("product name required")
	}
	if p.Description == "" {
		return errors.New("product description required")
	}
	if p.StorageUnit == enums.UNDEFINED {
		return errors.New("product storage unit cannot be undefined")
	}
	if p.UnitValue < 1.0 {
		return errors.New("product unit value cannot be less than 1.0")
	}

	return nil
}
