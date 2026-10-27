package main

import (
	"dz4/internal/auth/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		log.Println("Failed to open database")
		panic(err)
	}
	err = db.AutoMigrate(models.Product{})
	if err != nil {
		panic(err)
	}
}
