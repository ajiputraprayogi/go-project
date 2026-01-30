package main

import (
	"log"
	"os"

	"go-project/config"
	"go-project/models"
	"go-project/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env hanya untuk local
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{}, &models.Post{})

	r := routes.SetupRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // fallback buat local
	}

	r.Run(":" + port)
}
