package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var KEY string = "super_sneaky_key"

func GenerateToken(email string, userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userID": userID,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(KEY))
}