package config

import (
	"kentang/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	connectionString := "root:kentang@tcp(reviewhpdb:3306)/reviewhp?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}

func InitialMigration() {
	DB := Connect()
	DB.AutoMigrate(&models.User{}, &models.Phone{}, &models.Comment{})
}
