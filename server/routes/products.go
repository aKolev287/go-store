package routes

import (
	"go-store-server/db"
	"go-store-server/models"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)



func createProduct(ctx *gin.Context) {
	var product models.Product
	err := ctx.ShouldBindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	userID := ctx.GetUint("userID")
	product.UserID = userID

	product.ID = 0

	models.Save(&product)

	ctx.JSON(http.StatusCreated, gin.H{"message": "Product created!", "product": product})
}

func readProduct(ctx *gin.Context) {
	var product models.Product
	productId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Query failed"})
		return
	}

	models.Read(&product, uint(productId))

	ctx.JSON(http.StatusOK, product)
}

func readAllProducts(ctx *gin.Context) {
	var products []models.Product
	result := models.ReadAll(&products)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Query failed"})
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No records found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"products": products})
}

func updateProduct(ctx *gin.Context) {
	var product models.Product
	productId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid product ID"})
		return
	}

	models.Read(&product, uint(productId))

	userID := ctx.GetUint("userID")

	if product.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	var updatedProduct map[string]interface{}
	err = ctx.ShouldBindJSON(&updatedProduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	if len(updatedProduct) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "No fields to update"})
		return
	}

	if err := db.DB.First(&product, productId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}

	if err := db.DB.Model(&product).Updates(updatedProduct).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update product"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product successfully updated", "product": product})
}

func deleteProduct(ctx *gin.Context) {
	var product models.Product
	productId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid product ID"})
		return
	}

	models.Read(&product, uint(productId))

	userID := ctx.GetUint("userID")

	if product.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	models.Delete(&product, uint(productId))

	ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
