package usecases

import (
	"github.com/icezatoo/demo-go-api/pkg/domains"
	dto "github.com/icezatoo/demo-go-api/pkg/dto/user"
	"github.com/jinzhu/copier"
)

type userUseCase struct {
	repo domains.UserRepository
}

func NewUserUseCase(repo domains.UserRepository) *userUseCase {
	return &userUseCase{repo}
}

func (u *userUseCase) GetUsers() ([]*dto.UserReponse, error) {
	var usersDto = make([]*dto.UserReponse, 0)

	users, err := u.repo.GetUsers()

	copier.Copy(&usersDto, &users)

	return usersDto, err
}

func (u *userUseCase) GetUser(request *dto.RequestGetUser) (*dto.UserReponse, error) {
	var userDto *dto.UserReponse

	user, err := u.repo.GetUser(request)

	copier.Copy(&userDto, &user)

	return userDto, err
}

func (u *userUseCase) CreateUser(request *dto.CreateUserRequest) (*dto.UserReponse, error) {
	var userDto *dto.UserReponse

	user, err := u.repo.CreateUser(request)

	copier.Copy(&userDto, &user)

	return userDto, err
}

func (u *userUseCase) UpdateUser(request *dto.UpdateUserRequest) (*dto.UserReponse, error) {
	var userDto *dto.UserReponse

	user, err := u.repo.UpdateUser(request)

	copier.Copy(&userDto, &user)

	return userDto, err
}

func (u *userUseCase) DeleteUser(request *dto.RequestDeleteUser) error {

	err := u.repo.DeleteUser(request)

	return err
}
