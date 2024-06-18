package models

import (
	"go-store-server/db"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Price       uint
	Reviwes     int
	Description string
}

func Migrate() {
	db.DB.AutoMigrate(&Product{})
}

func Save(p *Product) {
	db.DB.Create(&p)
}

func Read(p *Product, ID uint) {
	db.DB.First(&p, ID)
}

func ReadAll(p *[]Product) *gorm.DB {
	result := db.DB.Find(&p)
	return result
}

func Delete(p *Product, ID uint) {
	db.DB.Delete(&p, ID)
}
