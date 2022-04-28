package models

import (
	"gorm.io/gorm"
)

type Auth struct {
	gorm.Model
	Name  string
	Email string
	Token string
}
