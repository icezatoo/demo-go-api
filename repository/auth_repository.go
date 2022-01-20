package repository

import (
	model "github.com/icezatoo/demo-go-api/models"
	"github.com/icezatoo/demo-go-api/utils"
	"gorm.io/gorm"
)

type AuthRepository interface {
	LoginRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

func NewRepositoryAuth(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) LoginRepository(input *model.EntityUsers) (*model.EntityUsers, string) {

	var users model.EntityUsers
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	users.Email = input.Email
	users.Password = input.Password

	hasUser := db.Debug().Select("*").Where("email = ?", input.Email).Find(&users)

	if hasUser.RowsAffected < 1 {
		errorCode <- "LOGIN_NOT_FOUND_404"
		return &users, <-errorCode
	}

	if !users.Enabled {
		errorCode <- "LOGIN_NOT_ACTIVE_403"
		return &users, <-errorCode
	}

	comparePassword := utils.ComparePassword(users.Password, input.Password)

	if comparePassword != nil {
		errorCode <- "LOGIN_WRONG_PASSWORD_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
