package main

import (
	"datingapp/internal/configs"
	"datingapp/internal/handlers"
	"datingapp/internal/repositories"
	"datingapp/internal/routes"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := configs.InitDB()
	userRepo := repositories.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)
	router := routes.SetupRoutes(userHandler)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
		return
	}
}
