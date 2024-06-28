package middleware

import (
	"go-store-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	userID, err := utils.VerifyToken(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid token."})
		return
	}

	ctx.Set("userID", userID)
	ctx.Next()
}
