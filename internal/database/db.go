package database

import (
	"final-project-01/internal/config"
	"final-project-01/internal/domain"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	err error
)

func StartDB() {
	config := config.GetDBConfig()

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	db.Debug().AutoMigrate(domain.Todo{})
}

func GetDB() *gorm.DB {
	return db
}