package repository

import (
	dto "github.com/icezatoo/demo-go-api/pkg/dto/user"
	customError "github.com/icezatoo/demo-go-api/pkg/errors"
	model "github.com/icezatoo/demo-go-api/pkg/model"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type UserRepository interface {
	GetUsers() ([]*model.User, error)
	GetUser(request *dto.RequestGetUser) (*model.User, error)
	CreateUser(request *dto.CreateUserRequest) (*model.User, error)
	UpdateUser(request *dto.UpdateUserRequest) (*model.User, error)
	DeleteUser(request *dto.RequestDeleteUser) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &repository{db: db}
}

func (repo *repository) GetUsers() ([]*model.User, error) {
	var users []*model.User

	err := repo.db.Find(&users).Error

	return users, err
}

func (repo *repository) GetUser(request *dto.RequestGetUser) (*model.User, error) {
	var user model.User

	err := repo.db.First(&user, "id = ?", request.ID).Error

	return &user, err
}

func (repo *repository) CreateUser(request *dto.CreateUserRequest) (*model.User, error) {
	var user model.User

	result := repo.db.Select("*").Where("email = ?", request.Email).Find(&user)

	if result.RowsAffected > 0 {
		return &user, customError.AlreadyExists("Email already exists")
	}

	user.FullName = request.FullName
	user.Email = request.Email
	user.LastName = request.LastName
	user.Password = request.Password
	user.Enabled = request.Enabled

	err := repo.db.Create(&user).Error

	return &user, err
}

func (repo *repository) UpdateUser(request *dto.UpdateUserRequest) (*model.User, error) {
	var user model.User
	result := repo.db.Select("*").Where("id = ?", request.ID).Find(&user)

	if result.RowsAffected < 1 {
		return &user, customError.NotFound("User not found")
	}

	user.FullName = request.FullName
	user.Email = request.Email
	user.LastName = request.LastName
	user.Enabled = request.Enabled

	err := repo.db.Updates(&user).Error

	return &user, err

}

func (repo *repository) DeleteUser(request *dto.RequestDeleteUser) error {
	var user model.User

	result := repo.db.Select("id").Where("id = ?", request.ID).First(&user)

	if result.RowsAffected < 1 {
		return customError.NotFound("User not found")
	}

	err := repo.db.Unscoped().Delete(&user, "id = ?", request.ID).Error

	return err
}
