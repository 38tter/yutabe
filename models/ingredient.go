package models

import "github.com/jinzhu/gorm"

type Ingredient struct {
	gorm.Model
	RecipeID uint
	Name     string `gorm:"varchar(50);not null"`
}

type IngredientJson struct {
	gorm.Model
	Name string `json:"name"`
}
