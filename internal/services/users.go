package services

import (
	"app/internal/models"
	"app/internal/repositories"
)

type UserService struct {
	repo repositories.Repository
}

func (s *UserService) Create(u models.User) (models.User, error) {
	if err := u.ValidateUser(); err != nil {
		return models.User{}, err
	}

	u, err := s.repo.Users.Create(u)
	if err != nil {
		return models.User{}, err
	}

	return u, nil
}

func (s *UserService) GetAll() ([]models.User, error) {

	users, err := s.repo.Users.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}
