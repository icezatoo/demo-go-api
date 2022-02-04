package repositories

import (
	"errors"

	"github.com/icezatoo/demo-go-api/pkg/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() ([]*entities.EntityUsers, error)
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
