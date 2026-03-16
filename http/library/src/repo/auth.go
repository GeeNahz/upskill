package repo

import (
	"http/library/src/models"

	"gorm.io/gorm"
)

func CreateUser(user *models.User, db *gorm.DB) *gorm.DB {
	return db.Create(user)

}

func GetUserByUsername(username string, db *gorm.DB, dbUser *models.User) *gorm.DB {
	return db.Where("username = ?", username).First(dbUser)
}
