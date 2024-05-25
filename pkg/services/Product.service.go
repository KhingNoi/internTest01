package services

import (
	"internTest01/pkg/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProductList(context *gin.Context, db *gorm.DB) {
	products, err := repositories.FindProductList(db)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, products)
}
