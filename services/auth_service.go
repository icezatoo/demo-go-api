package services

import (
	model "github.com/icezatoo/demo-go-api/models"
	"github.com/icezatoo/demo-go-api/repository"
	requests "github.com/icezatoo/demo-go-api/requests"
)

type AuthService interface {
	LoginService(input *requests.InputLogin) (*model.EntityUsers, string)
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(repository repository.AuthRepository) *authService {
	return &authService{repository: repository}
}

func (s *authService) LoginService(input *requests.InputLogin) (*model.EntityUsers, string) {

	user := model.EntityUsers{
		Email:    input.Email,
		Password: input.Password,
	}

	resultLogin, errLogin := s.repository.LoginRepository(&user)

	return resultLogin, errLogin
}
