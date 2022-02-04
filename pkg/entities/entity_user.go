package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/icezatoo/demo-go-api/pkg/utils"
	"gorm.io/gorm"
)

type EntityUsers struct {
	gorm.Model
	ID       string `gorm:"primary_key" json:"id"`
	Enabled  bool   `gorm:"type:bool;default:false" json:"enable"`
	FullName string `gorm:"full_name;varchar(255);" json:"fullName"`
	LastName string `gorm:"last_name;varchar(255)" json:"lastName"`
	Email    string `gorm:"email;varchar(255);unique;not null" json:"email"`
	Password string `gorm:"password;varchar(255);unique;not null" json:"password"`
}

func (EntityUsers) TableName() string {
	return "users"
}

func (entity *EntityUsers) BeforeCreate(tx *gorm.DB) (err error) {
	entity.ID = uuid.New().String()
	entity.Password = utils.HashPassword(entity.Password)
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *EntityUsers) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
