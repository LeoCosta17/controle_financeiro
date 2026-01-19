package models

import "errors"

type Costumer struct {
	ID        uint64 `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Telephone string `json:"telephone,omitempty"`
	Document  string `json:"document,omitempty"`
}

func (c *Costumer) ValidateCostumer() error {
	if err := validateCostumerData(c); err != nil {
		return err
	}

	return nil
}

func validateCostumerData(c *Costumer) error {
	if c.Name == "" {
		return errors.New("costumer name required")
	}
	return nil
}
