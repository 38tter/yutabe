package models

import "github.com/jinzhu/gorm"

type InstructionText struct {
	gorm.Model
	InstructionID uint
	Text          string `gorm:"varchar(255);not null"`
}
