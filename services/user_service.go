package services

import (
	model "github.com/icezatoo/demo-go-api/models"
	"github.com/icezatoo/demo-go-api/repository"
	requests "github.com/icezatoo/demo-go-api/requests"
)

type UserService interface {
	CreateUser(input *requests.InputRegisterUser) (*model.EntityUsers, string)
	UpdateUser(input *requests.InputUpdateUser) (*model.EntityUsers, string)
	DeleteUser(input *requests.InputDeleteUser) (*model.EntityUsers, string)
	GetUser(input *requests.InputGetUser) (*model.EntityUsers, string)
	GetUserList() (*[]model.EntityUsers, string)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository: repository}
}

func (s *userService) CreateUser(input *requests.InputRegisterUser) (*model.EntityUsers, string) {
	user := model.EntityUsers{
		FullName: input.FullName,
		Email:    input.Email,
		Password: input.Password,
		Enabled:  input.Enabled,
	}
	resultRegister, err := s.repository.CreateUserRepository(&user)
	return resultRegister, err
}

func (s *userService) UpdateUser(input *requests.InputUpdateUser) (*model.EntityUsers, string) {
	user := model.EntityUsers{
		FullName: input.FullName,
		Email:    input.Email,
		Enabled:  input.Enabled,
	}
	resultUpdate, err := s.repository.UpdateUserRepository(&user)
	return resultUpdate, err
}

func (s *userService) DeleteUser(input *requests.InputDeleteUser) (*model.EntityUsers, string) {
	user := model.EntityUsers{
		ID: input.ID,
	}
	resultDelete, err := s.repository.DeleteUserRepository(&user)
	return resultDelete, err
}

func (s *userService) GetUser(input *requests.InputGetUser) (*model.EntityUsers, string) {
	user := model.EntityUsers{
		ID: input.ID,
	}
	resultGet, err := s.repository.GetUserRepository(&user)
	return resultGet, err
}

func (s *userService) GetUserList() (*[]model.EntityUsers, string) {
	resultList, err := s.repository.GetUserListRepository()
	return resultList, err
}
