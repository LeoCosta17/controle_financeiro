package models

import "errors"

type Supplier struct {
	ID        uint64 `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Telephone string `json:"telephone,omitempty"`
	Document  string `json:"document,omitempty"`
}

func (s *Supplier) ValidateSupplier() error {
	if err := validateSupplierData(s); err != nil {
		return err
	}

	return nil
}

func validateSupplierData(s *Supplier) error {
	if s.Name == "" {
		return errors.New("supplier name required")
	}
	return nil
}
