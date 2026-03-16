package database

import (
	"fmt"
	"os"

	"github.com/geenahz/recipe/src/models"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBInterface interface {
	ConnectDB() *DB
	Migrate()
}

type DB struct {
	DB *gorm.DB
}

func ConnectDB() *DB {
	_ = godotenv.Load()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	passwd := os.Getenv("DB_PASSWD")
	dbName := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || passwd == "" || dbName == "" {
		log.Fatal("Missing environment variables")
	}

	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Africa/Lagos", host, user, passwd, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	return &DB{DB: db}
}

func (db *DB) Migrate() {
	db.DB.AutoMigrate(
		&models.Recipe{},
		&models.User{},
		&models.Instructions{},
		&models.Ingredients{},
	)

	tx := db.DB.Begin()
	tx.Commit()
}
