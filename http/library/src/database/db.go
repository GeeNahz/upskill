package database

import (
	"http/library/src/models"

	"github.com/gofiber/fiber/v2/log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("storage.db"))

	if err != nil {
		log.Fatal("Failed to connect to the database")
	}

	db.AutoMigrate(&models.User{}, &models.Book{})

	return db
}
