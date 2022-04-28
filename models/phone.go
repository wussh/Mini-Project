package models

import (
	"gorm.io/gorm"
)

type Phone struct {
	gorm.Model
	Price      string
	Design     string
	Display    string
	Perfomance string
	Cameras    string
	Audio      string
	Battery    string
	Features   string
}
