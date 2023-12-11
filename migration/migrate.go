package main

import (
	"go_crud/initializers"
	model "go_crud/models"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("failed to perform auto-migration:", err)
	} else {
		log.Print("auto-migration completed successfully")
	}
}
