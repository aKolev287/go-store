package routes

import (
	"go-store-server/models"
	"go-store-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)



func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Faild to create user!"})
		return
	}

	user.ID = 0

	user.SaveUser()
	ctx.JSON(http.StatusCreated, gin.H{"message": "Successfully created a new user!"})
}

func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

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


func fetchUser(ctx *gin.Context) {
	var user models.User
	cookie, err := ctx.Request.Cookie("token")
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Invalid or missing token!"})
		return
	}

	if cookie.Value == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	userID, err := utils.VerifyToken(cookie.Value)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid token."})
		return
	}

	user.ReadUser(userID)

	ctx.JSON(http.StatusOK, user)
	
}

