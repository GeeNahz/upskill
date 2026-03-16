package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	PrepTime    int    `json:"prepTime"`
	CookTime    int    `json:"cookTime"`
	Notes       string `json:"notes"`

	Ingredients  []Ingredients  `gorm:"constraint:OnUpdate:CAsCADE,OnDelete:CASCADE" json:"ingredients"`
	Instructions []Instructions `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"instructions"`

	UserID uint `json:"user_id"`
	Chef   User `gorm:"foreignKey:UserID" json:"chef"`

	CreatedBy *uint
	UpdatedBy *uint
}

type Ingredients struct {
	gorm.Model
	Name     string `json:"name"`
	Quantity string `json:"quantity"`

	RecipeID uint `json:"recipe_id"`
}

type Instructions struct {
	gorm.Model
	Step  string `json:"step"`
	Order int    `json:"order"`

	RecipeID uint `json:"recipe_id"`
}

// type IRecipe interface {
// 	GetAll() []Recipe
// 	GetByID(id int) Recipe
// 	Create(recipe Recipe) Recipe
// 	Update(id int, recipe Recipe) Recipe
// 	Delete(id int) bool
// }
