package routes

import (
	"go-store-server/models"
	"go-store-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var user models.User

func signup(ctx *gin.Context) {
	err := ctx.ShouldBindBodyWithJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Faild to create user!"})
		return
	}
	user.SaveUser()
	ctx.JSON(http.StatusCreated, gin.H{"message": "Successfully created a new user!"})
}

func login(ctx *gin.Context) {
	err := ctx.ShouldBindBodyWithJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid credentials"})
		return
	}

	err = user.ValidateUser(user.Email, user.Password)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.Username, user.ID)

	ctx.SetCookie("token", token, 7200, "/", "localhost", false, true)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to validate credentials"})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Log in successful", "token": token})
}
