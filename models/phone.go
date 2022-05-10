package models

import (
	"gorm.io/gorm"
)

type Phone struct {
	gorm.Model
	Name       string
	Price      string
	Design     string
	Display    string
	Perfomance string
	Cameras    string
	Audio      string
	Battery    string
	Features   string
	Comment    []Comment `gorm:"foreignKey:PhoneRefer"`
}
