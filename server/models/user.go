package models

import (
	"errors"
	"fmt"
	"go-store-server/db"
	"go-store-server/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
	role     string
}

func (u *User) SaveUser() {
	hashedPass, err := utils.HashPassword(u.Password)
	u.Password = hashedPass
	if err != nil {
		return
	}

	db.DB.Create(&u)
}

func (u *User) ValidateUser(email string, password string) error {

	db.DB.Select("ID", "email", "password").Where("email = ?", email).Find(&u)

	fmt.Println(password)
	fmt.Println(u.Password)

	err := utils.ComparePassword(password, u.Password)

	if !err {
		return errors.New("password does not match")
	}

	return nil
}