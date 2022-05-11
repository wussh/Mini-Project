package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserNumber  string
	PhoneNumber string
	Context     string `json:"context" form:"context"`
}
