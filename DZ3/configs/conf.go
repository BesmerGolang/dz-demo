package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EmailConfig struct {
	Email    string
	Password string
	Address  string
}

func LoadConfig() *EmailConfig {
	err := godotenv.Load() //функция загружающая конфиг из пакета годотенв
	if err != nil {
		log.Printf("error loading env file: %v", err)
	}
	return &EmailConfig{
		Email:    os.Getenv("EMAIL"),
		Password: os.Getenv("PASSWORD"),
		Address:  os.Getenv("ADDRESS"),
	}
}
