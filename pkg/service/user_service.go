package service

import (
	dto "github.com/icezatoo/demo-go-api/pkg/dto/user"
	"github.com/icezatoo/demo-go-api/pkg/repository"
	"github.com/jinzhu/copier"
)

type userService struct {
	repo repository.UserRepository
}

type UserService interface {
	GetUsers() ([]*dto.UserReponse, error)
	GetUser(request *dto.RequestGetUser) (*dto.UserReponse, error)
	CreateUser(request *dto.CreateUserRequest) (*dto.UserReponse, error)
	UpdateUser(request *dto.UpdateUserRequest) (*dto.UserReponse, error)
	DeleteUser(request *dto.RequestDeleteUser) error
}

func NewUserUseCase(repo repository.UserRepository) *userService {
	return &userService{repo}
}

func (u *userService) GetUsers() ([]*dto.UserReponse, error) {
	var usersDto = make([]*dto.UserReponse, 0)

	users, err := u.repo.GetUsers()

	copier.Copy(&usersDto, &users)

	return usersDto, err
}

func (u *userService) GetUser(request *dto.RequestGetUser) (*dto.UserReponse, error) {
	var userDto *dto.UserReponse

	user, err := u.repo.GetUser(request)

	copier.Copy(&userDto, &user)

	return userDto, err
}

func (u *userService) CreateUser(request *dto.CreateUserRequest) (*dto.UserReponse, error) {
	var userDto *dto.UserReponse

	user, err := u.repo.CreateUser(request)

	copier.Copy(&userDto, &user)

	return userDto, err
}

func (u *userService) UpdateUser(request *dto.UpdateUserRequest) (*dto.UserReponse, error) {
	var userDto *dto.UserReponse

	user, err := u.repo.UpdateUser(request)

	copier.Copy(&userDto, &user)

	return userDto, err
}

func (u *userService) DeleteUser(request *dto.RequestDeleteUser) error {

	err := u.repo.DeleteUser(request)

	return err
}
