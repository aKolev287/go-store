package routes

import (
	"go-store-server/db"
	"go-store-server/middleware"
	"go-store-server/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Migrate() {
	db.DB.AutoMigrate(&models.Product{})
	db.DB.AutoMigrate(&models.User{})
}

func ServerRouter(r *gin.Engine) {
	
	// Allow all for debug
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowCredentials = true


	r.Use(cors.New(config))


	r.GET("/product/:id/", readProduct)
	r.GET("/products/", readAllProducts)

	authenticated := r.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/product/", createProduct)
	authenticated.PUT("/product/:id/", updateProduct)
	authenticated.DELETE("/product/:id/", deleteProduct)

	r.GET("/user/", fetchUser)

	r.POST("/user/signup/", signup)
	r.POST("/user/login/", login)

}
