package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var KEY string = "super_sneaky_key"

func GenerateToken(email string, userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(KEY))
}

func VerifyToken(token string) (uint, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {

		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(KEY), nil
	})

	if err != nil {
		return 0, errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token")
	}
	userID := claims["userID"].(float64)

	return uint(userID), nil
}
