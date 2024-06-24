package routes

import (
	"go-store-server/db"
	"go-store-server/models"

	"github.com/gin-gonic/gin"
)

func Migrate() {
	db.DB.AutoMigrate(&models.Product{})
	db.DB.AutoMigrate(&models.User{})
}

func ServerRouter(r *gin.Engine) {
	r.POST("product/", createProduct)
	r.GET("product/:id/", readProduct)
	r.GET("products/", readAllProducts)
	r.PUT("product/:id/", updateProduct)
	r.DELETE("product/:id/", deleteProduct)

	r.POST("user/signup/", signup)
	r.POST("user/login/", login)
}
