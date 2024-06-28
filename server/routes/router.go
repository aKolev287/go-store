package routes

import (
	"go-store-server/db"
	"go-store-server/middleware"
	"go-store-server/models"

	"github.com/gin-gonic/gin"
)

func Migrate() {
	db.DB.AutoMigrate(&models.Product{})
	db.DB.AutoMigrate(&models.User{})
}

func ServerRouter(r *gin.Engine) {
	
	r.GET("/product/:id/", readProduct)
	r.GET("/products/", readAllProducts)

	authenticated := r.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/product/", createProduct)
	authenticated.PUT("/product/:id/", updateProduct)
	authenticated.DELETE("/product/:id/", deleteProduct)

	r.POST("/user/signup/", signup)
	r.POST("/user/login/", login)
}
