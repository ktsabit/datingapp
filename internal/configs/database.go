package configs

import (
	"datingapp/internal/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitDB() *gorm.DB {

	dbHost := os.Getenv("DB_HOST")

	log.Println(dbHost)

	dsn := fmt.Sprintf("host=%s user=admin password=admin123 dbname=datingapp port=5432 sslmode=disable", dbHost)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("Connected to database")

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	return db
}
