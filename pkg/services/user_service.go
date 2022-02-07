package services

import (
	dto "github.com/icezatoo/demo-go-api/pkg/dto/user"
	"github.com/icezatoo/demo-go-api/pkg/repositories"
)

type UserService interface {
	GetUsers() ([]*dto.UserReponse, error)
	GetUser(request *dto.RequestGetUser) (*dto.UserReponse, error)
	CreateUser(request *dto.CreateUserRequest) (*dto.UserReponse, error)
	DeleteUser(request *dto.RequestDeleteUser) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *userService {
	return &userService{repo}
}

func (u *userService) GetUsers() ([]*dto.UserReponse, error) {
	var usersDto = make([]*dto.UserReponse, 0)

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

func (u *userService) GetUser(request *dto.RequestGetUser) (*dto.UserReponse, error) {
	var userDto dto.UserReponse

	user, err := u.repo.GetUser(request)

	userDto.ID = user.ID
	userDto.Email = user.Email
	userDto.FullName = user.FullName
	userDto.LastName = user.LastName
	userDto.Enabled = user.Enabled

	return &userDto, err
}

func (u *userService) CreateUser(request *dto.CreateUserRequest) (*dto.UserReponse, error) {
	var userDto dto.UserReponse
	user, err := u.repo.CreateUser(request)

	userDto.ID = user.ID
	userDto.Email = user.Email
	userDto.FullName = user.FullName
	userDto.LastName = user.LastName
	userDto.Enabled = user.Enabled
	return &userDto, err
}

func (u *userService) DeleteUser(request *dto.RequestDeleteUser) error {
	return u.repo.DeleteUser(request)
}
