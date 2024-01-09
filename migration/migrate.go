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
	initializers.DB = initializers.DB.Debug()
	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	err := initializers.DB.AutoMigrate(&model.User{}, &model.Post{})
	if err != nil {
		log.Fatal("failed to perform auto-migration:", err)
	} else {
		log.Print("auto-migration completed successfully")
	}
}
