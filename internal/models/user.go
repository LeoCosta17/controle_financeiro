package models

import "errors"

type User struct {
	ID       uint64 `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (u *User) ValidateUser() error {
	if err := validateData(u); err != nil {
		return err
	}
	return nil
}

func validateData(u *User) error {
	if u.Name == "" {
		return errors.New("user name required")
	}
	if u.Email == "" {
		return errors.New("user email required")
	}
	if u.Password == "" {
		return errors.New("user password required")
	}
	return nil
}
