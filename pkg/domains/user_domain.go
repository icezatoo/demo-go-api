package domains

import (
	dto "github.com/icezatoo/demo-go-api/pkg/dto/user"
	"github.com/icezatoo/demo-go-api/pkg/entities"
)

type UserUseCase interface {
	GetUsers() ([]*dto.UserReponse, error)
	GetUser(request *dto.RequestGetUser) (*dto.UserReponse, error)
	CreateUser(request *dto.CreateUserRequest) (*dto.UserReponse, error)
	UpdateUser(request *dto.UpdateUserRequest) (*dto.UserReponse, error)
	DeleteUser(request *dto.RequestDeleteUser) error
}

type UserRepository interface {
	GetUsers() ([]*entities.EntityUsers, error)
	GetUser(request *dto.RequestGetUser) (*entities.EntityUsers, error)
	CreateUser(request *dto.CreateUserRequest) (*entities.EntityUsers, error)
	UpdateUser(request *dto.UpdateUserRequest) (*entities.EntityUsers, error)
	DeleteUser(request *dto.RequestDeleteUser) error
}
