package repository

import (
	model "github.com/icezatoo/demo-go-api/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUserRepository(input *model.EntityUsers) (*model.EntityUsers, string)
	DeleteUserRepository(input *model.EntityUsers) (*model.EntityUsers, string)
	UpdateUserRepository(input *model.EntityUsers) (*model.EntityUsers, string)
	GetUserListRepository() (*[]model.EntityUsers, string)
	GetUserRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

func NewRepositoryUser(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (repo *repository) UpdateUserRepository(input *model.EntityUsers) (*model.EntityUsers, string) {
	var user model.EntityUsers
	errorCode := make(chan string, 1)
	db := repo.db.Model(&user)

	user.ID = input.ID
	hasUser := db.Select("*").Where("id = ?", input.ID).Find(&user)

	if hasUser.RowsAffected < 1 {
		errorCode <- "UPDATE_USER_NOT_FOUND_404"
		return &user, <-errorCode
	}

	user.FullName = input.FullName
	user.Enabled = input.Enabled
	user.Email = input.Email

	updateUser := db.Select("id", "full_name", "enabled", "email", "created_at", "updated_at").Where("id = ?", input.ID).Updates(user)

	if updateUser.Error != nil {
		errorCode <- "UPDATE_USER_FAILED_403"
		return &user, <-errorCode
	} else {
		errorCode <- "nil"
	}
	return &user, <-errorCode
}

func (repo *repository) DeleteUserRepository(input *model.EntityUsers) (*model.EntityUsers, string) {
	var user model.EntityUsers
	errorCode := make(chan string, 1)
	db := repo.db.Model(&user)

	hasUser := db.Select("*").Where("id = ?", input.ID).Find(&user)

	if hasUser.RowsAffected < 1 {
		errorCode <- "DELETE_USER_NOT_FOUND_404"
		return &user, <-errorCode
	}

	deleteUser := db.Select("*").Where("id = ?", input.ID).Find(&user).Delete(&user)

	if deleteUser.Error != nil {
		errorCode <- "DELETE_USER_FAILED_403"
		return &user, <-errorCode
	} else {
		errorCode <- "nil"
	}
	return &user, <-errorCode
}

func (repo *repository) CreateUserRepository(input *model.EntityUsers) (*model.EntityUsers, string) {
	var user model.EntityUsers

	errorCode := make(chan string, 1)

	db := repo.db.Model(&user)

	hasUser := db.Select("*").Where("email = ?", input.Email).Find(&user)
	if hasUser.RowsAffected > 0 {
		errorCode <- "USER_CONFLICT_409"
		return &user, <-errorCode
	}

	user.FullName = input.FullName
	user.Enabled = input.Enabled
	user.Email = input.Email
	user.Password = input.Password

	addUser := db.Create(&user)

	if addUser.Error != nil {
		errorCode <- "CREATE_USER_FAILED_403"
		return &user, <-errorCode
	} else {
		errorCode <- "nil"
	}
	return &user, <-errorCode
}

func (repo *repository) GetUserListRepository() (*[]model.EntityUsers, string) {
	var users []model.EntityUsers
	db := repo.db.Model(&users)
	errorCode := make(chan string, 1)

	resultsUsers := db.Debug().Select("*").Find(&users)

	if resultsUsers.Error != nil {
		errorCode <- "RESULTS_USER_NOT_FOUND_404"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}

func (repo *repository) GetUserRepository(input *model.EntityUsers) (*model.EntityUsers, string) {
	var user model.EntityUsers
	errorCode := make(chan string, 1)
	db := repo.db.Model(&user)

	resultUser := db.Select("*").Where("id = ?", input.ID).Find(&user)

	if resultUser.RowsAffected < 1 {
		errorCode <- "RESULT_USER_NOT_FOUND_404"
		return &user, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &user, <-errorCode
}
