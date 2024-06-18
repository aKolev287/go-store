package routes

import (
	"github.com/gin-gonic/gin"
)

func ServerRouter(r *gin.Engine) {
	r.POST("product/", createProduct)
	r.GET("product/:id/", readProduct)
	r.GET("products/", readAllProducts)
	r.PUT("product/:id/", updateProduct)
	r.DELETE("product/:id/", deleteProduct)
}
