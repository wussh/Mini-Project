package main

import (
	"kentang/config"
	"kentang/routes"
)

func main() {
	config.Connect()
	config.InitialMigration()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
