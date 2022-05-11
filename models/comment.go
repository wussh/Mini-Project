package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId  string
	PhoneId string
	Context string `json:"context" form:"context"`
}
