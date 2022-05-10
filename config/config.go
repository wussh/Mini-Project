package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	connectionString := "root:kentang@tcp(127.0.0.1:3306)/merdeka?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}
