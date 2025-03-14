package main

import (
	"datingapp/internal/configs"
	"datingapp/internal/handlers"
	"datingapp/internal/repositories"
	"datingapp/internal/routes"
	"datingapp/internal/services"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtService := services.NewJWTService(services.JWTConfig{
		AccessTokenSecret:  os.Getenv("ACCESS_TOKEN_SECRET"),
		RefreshTokenSecret: os.Getenv("REFRESH_TOKEN_SECRET"),
		AccessTokenTTL:     15 * time.Minute,
		RefreshTokenTTL:    30 * 24 * time.Hour,
	})

	db := configs.InitDB()
	rdb := configs.InitRedis()
	userRepo := repositories.NewUserRepository(db)
	swipeRepo := repositories.NewSwipeRepository(db, rdb)
	userService := services.NewUserService(userRepo)
	swipeService := services.NewSwipeService(swipeRepo, userRepo)
	userHandler := handlers.NewUserHandler(userService)
	swipeHandler := handlers.NewSwipeHandler(swipeService)
	authHandler := handlers.NewAuthHandler(userRepo, jwtService)

	router := routes.SetupRoutes(userHandler, authHandler, jwtService, swipeHandler)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
		return
	}
}
