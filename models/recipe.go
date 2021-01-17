package models

import "github.com/jinzhu/gorm"

type Recipe struct {
	gorm.Model
	Name        string `gorm:"varchar(100);not null"`
	ImageUrl    string `gorm:"varchar(100);not null"`
	Description string `gorm:"varchar(100);not null"`
	TastyID     int    `gorm:"int;not null"`
	Instruction *Instruction
	ingredients []*Ingredient
	//rate         float64 `gorm:"varchar(100);not null"`
	//utritions    []*Nutrition
}

type RecipeJson struct {
	Name        string            `json:"name"`
	ImageUrl    string            `json:"image_url"`
	Description string            `json:"description"`
	TastyID     int               `json:"tasty_id"`
	Instruction *InstructionJson  `json:"instruction"`
	Ingredients []*IngredientJson `json:"ingredients"`
	//rate         float64 `gorm:"varchar(100);not null"`
	//utritions    []*Nutrition
}
