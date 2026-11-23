package db

import (
	"dz4/configs"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(conf.Db.Dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to open database")
		panic(err)
	}
	return &Db{db}
}
