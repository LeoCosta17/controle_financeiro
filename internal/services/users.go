package services

import (
	"app/internal/models"
)

type userService struct {
}

func NewUserService() userService {
	return userService{}
}

func (s *userService) CreateUser(u models.User) (models.User, error) {
	return models.User{}, nil
}
