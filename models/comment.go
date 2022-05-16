package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID  int
	User    User
	PhoneID int
	Phone   Phone
	Context string `json:"context" form:"context"`
}
