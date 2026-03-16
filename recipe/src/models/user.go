package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string   `gorm:"uniqueIndex" json:"email"`
	Name     string   `gorm:"size:100" json:"name"`
	Password string   `json:"-"`
	Recipes  []Recipe `gorm:"foreignKey:UserId"`

	CreatedBy *uint
	UpdatedBy *uint
}
