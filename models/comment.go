package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserRefer  uint
	PhoneRefer uint
	Context    string
}
