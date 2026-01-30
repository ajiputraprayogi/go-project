package main

import (
	"go-project/config"
	"go-project/models"
	"go-project/routes"
)

func main() {
	config.ConnectDB()	
	config.DB.AutoMigrate(&models.User{}, &models.Post{})

	r := routes.SetupRoutes()
	r.Run(":8000")
}
