package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Comment  []Comment `gorm:"foreignKey:UserRefer"`
}
