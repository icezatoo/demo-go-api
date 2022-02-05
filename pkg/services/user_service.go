package services

import (
	dto "github.com/icezatoo/demo-go-api/pkg/dto/user"
	"github.com/icezatoo/demo-go-api/pkg/repositories"
)

type UserService interface {
	GetUsers() ([]*dto.UserReponse, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *userService {
	return &userService{repo}
}

func (u *userService) GetUsers() ([]*dto.UserReponse, error) {
	var usersDto []*dto.UserReponse
	users, err := u.repo.GetUsers()

	for _, user := range users {
		usersDto = append(usersDto, &dto.UserReponse{
			ID:       user.ID,
			FullName: user.FullName,
			LastName: user.LastName,
			Email:    user.Email,
			Enabled:  user.Enabled,
		})
	}

	return usersDto, err
}
