package db

import (
	"github.com/icezatoo/demo-go-api/pkg/config"
	model "github.com/icezatoo/demo-go-api/pkg/model"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection(config *config.Config) *gorm.DB {

	db, err := gorm.Open(postgres.Open(config.DbURL), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
	}

	if config.Environment != "production" {
		logrus.Info("Connection to Database Successfully")
	}

	err = db.AutoMigrate(
		&model.User{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db

}
