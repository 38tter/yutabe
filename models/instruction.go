package models

import "github.com/jinzhu/gorm"

type Instruction struct {
	gorm.Model
	RecipeID uint
	Texts    []InstructionText
}

type InstructionJson struct {
	Texts []string `json:"texts"`
}
