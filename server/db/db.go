package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(keys string) {
	var err error

	DB, err = gorm.Open(postgres.Open(keys), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connected")
}
