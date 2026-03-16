package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"todo/src/models"
)

type Database struct {
	Db *gorm.DB
}

// var DB *gorm.DB

func ConnectDB() *Database {
	_ = godotenv.Load()

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	userName := os.Getenv("USERNAME")
	dbName := os.Getenv("DB_NAME")
	passwd := os.Getenv("PASSWORD")

	if host == "" || port == "" || userName == "" || dbName == "" || passwd == "" {
		log.Fatal("Missing required db credentials")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, userName, passwd, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	fmt.Println("Connected to database successfully")
	return &Database{Db: db}
}

func (DB *Database) MakeMigrations() {
	modelList := []any{
		&models.User{},
		&models.Todo{},
	}

	for _, model := range modelList {
		DB.Db.AutoMigrate(model)
	}

	tx := DB.Db.Begin()

	tx.Commit()
}
