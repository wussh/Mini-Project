package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	PhoneID int
	Phone   Phone
	Context string `json:"context" form:"context"`
}
