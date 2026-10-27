package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Images      pq.StringArray `gorm:"type:text[]"`
}
type Order struct {
	gorm.Model
	ProductID uint
	UserName  string
	Email     string
}
