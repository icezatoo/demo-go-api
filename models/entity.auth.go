package model

import (
	"time"

	"github.com/google/uuid"
	util "github.com/icezatoo/demo-go-api/utils"
	"gorm.io/gorm"
)

type EntityUsers struct {
	ID        string    `json:"id";gorm:"primaryKey;"`
	FullName  string    `json:"fullName";gorm:"type:varchar(255);unique;not null"`
	Email     string    `json:"email";gorm:"type:varchar(255);unique;not null"`
	Password  string    `json:"password";gorm:"type:varchar(255);not null"`
	Enabled   bool      `json:"enable";gorm:"type:bool;default:false"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updateAt"`
}

func (entity *EntityUsers) BeforeCreate(tx *gorm.DB) (err error) {
	entity.ID = uuid.New().String()
	entity.Password = util.HashPassword(entity.Password)
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *EntityUsers) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
