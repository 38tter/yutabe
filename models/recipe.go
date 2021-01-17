package models

import "github.com/jinzhu/gorm"

type Recipe struct {
	gorm.Model
	Name        string `gorm:"varchar(100);not null"`
	ImageUrl    string `gorm:"varchar(100);not null"`
	Description string `gorm:"varchar(100);not null"`
	Instruction *Instruction
	//rate         float64 `gorm:"varchar(100);not null"`
	//ingredients  []*Ingredient
	//utritions    []*Nutrition
}

type RecipeJson struct {
	Name        string           `json:"name"`
	ImageUrl    string           `json:"image_url"`
	Description string           `json:"description"`
	Instruction *InstructionJson `json:"instruction"`
	//rate         float64 `gorm:"varchar(100);not null"`
	//ingredients  []*Ingredient
	//utritions    []*Nutrition
}
