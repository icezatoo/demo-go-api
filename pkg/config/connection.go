package config

import (
	"os"

	"github.com/icezatoo/demo-go-api/pkg/entities"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	var databaseURI string

	if os.Getenv("GO_ENV") != "production" {
		databaseURI = os.Getenv("DATABASE_URI_DEV")
	} else {
		databaseURI = os.Getenv("DATABASE_URI_PROD")
	}

	db, err := gorm.Open(postgres.Open(databaseURI), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
	}

	if os.Getenv("GO_ENV") != "production" {
		logrus.Info("Connection to Database Successfully")
	}

	err = db.AutoMigrate(
		&entities.EntityUsers{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db

}
