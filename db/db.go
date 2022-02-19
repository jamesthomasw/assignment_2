package db

import (
	"log"

	"github.com/jamesthomasw/assignment_2/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:example@localhost:5432/orders_by"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	db.AutoMigrate(&models.Order{}, &models.Item{})

	return db
}