package services

import (
	"github.com/icezatoo/demo-go-api/pkg/entities"
	"github.com/icezatoo/demo-go-api/pkg/repositories"
)

type UserService interface {
	GetUsers() ([]*entities.EntityUsers, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *userService {
	return &userService{repo}
}

func (u *userService) GetUsers() ([]*entities.EntityUsers, error) {
	return u.repo.GetUsers()
}
