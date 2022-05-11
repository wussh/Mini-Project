package models

import (
	"gorm.io/gorm"
)

type Phone struct {
	gorm.Model
	Name       string    `json:"name" form:"name"`
	Price      string    `json:"price" form:"price"`
	Design     string    `json:"design" form:"design"`
	Display    string    `json:"display" form:"display"`
	Perfomance string    `json:"perfomance" form:"perfomance"`
	Cameras    string    `json:"cameras" form:"cameras"`
	Audio      string    `json:"audio" form:"audio"`
	Battery    string    `json:"battery" form:"battery"`
	Features   string    `json:"features" form:"features"`
	Comment    []Comment `gorm:"foreignKey:PhoneId;references:Name"`
}
