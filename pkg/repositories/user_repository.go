package repositories

import (
	"errors"

	dto "github.com/icezatoo/demo-go-api/pkg/dto/user"
	"github.com/icezatoo/demo-go-api/pkg/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() ([]*entities.EntityUsers, error)
	GetUser(request *dto.RequestGetUser) (*entities.EntityUsers, error)
	CreateUser(request *dto.CreateUserRequest) (*entities.EntityUsers, error)
	UpdateUser(request *dto.UpdateUserRequest) (*entities.EntityUsers, error)
	DeleteUser(request *dto.RequestDeleteUser) error
}

func NewRepositoryUser(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (repo *repository) GetUsers() ([]*entities.EntityUsers, error) {
	var users []*entities.EntityUsers

	err := repo.db.Find(&users).Error

	if err != nil {
		return nil, errors.New("GET_USERS_ERROR")
	}

	return users, nil
}

func (repo *repository) GetUser(request *dto.RequestGetUser) (*entities.EntityUsers, error) {
	var user entities.EntityUsers

	err := repo.db.First(&user, "id = ?", request.ID).Error

	return &user, err
}

func (repo *repository) CreateUser(request *dto.CreateUserRequest) (*entities.EntityUsers, error) {
	var user entities.EntityUsers

	result := repo.db.Select("*").Where("email = ?", request.Email).Find(&user)

	if result.RowsAffected > 0 {
		return &user, errors.New("USER_CONFLICT_409")
	}

	user.FullName = request.FullName
	user.Email = request.Email
	user.LastName = request.LastName
	user.Password = request.Password
	user.Enabled = request.Enabled

	err := repo.db.Create(&user).Error

	return &user, err
}

func (repo *repository) UpdateUser(request *dto.UpdateUserRequest) (*entities.EntityUsers, error) {
	var user entities.EntityUsers
	result := repo.db.Select("*").Where("id = ?", request.ID).Find(&user)

	if result.RowsAffected < 1 {
		return &user, errors.New("USER_NOT_FOUND_404")
	}

	user.FullName = request.FullName
	user.Email = request.Email
	user.LastName = request.LastName
	user.Enabled = request.Enabled

	err := repo.db.Updates(&user).Error

	return &user, err

}

func (repo *repository) DeleteUser(request *dto.RequestDeleteUser) error {
	var user entities.EntityUsers

	result := repo.db.Select("id").Where("id = ?", request.ID).First(&user)

	if result.RowsAffected < 1 {
		return errors.New("USER_NOT_FOUND_404")
	}

	err := repo.db.Unscoped().Delete(&user, "id = ?", request.ID).Error

	return err
}
