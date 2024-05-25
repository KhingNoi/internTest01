package services

import (
	"internTest01/pkg/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProductList(context *gin.Context, db *gorm.DB) {
	products, err := repositories.FindProductList(db)
	if len(products) == 0 {
		context.JSON(http.StatusNoContent, products)
	}
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, products)
}

func GetProductById(context *gin.Context, db *gorm.DB) {
	id := context.Param("id")
	if _, strconvErr := strconv.Atoi(id); strconvErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	product, err := repositories.FindProductById(db, id)
	if err != nil {
		if err.Error() == "record not found" {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	context.JSON(http.StatusOK, product)
}
