package services

import (
	"fmt"
	"internTest01/pkg/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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

func GetProductListWithPagination(context *gin.Context, db *gorm.DB) {
	fmt.Println("trigger")
	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(context.DefaultQuery("size", "12"))
	products, total, err := repositories.FindProductListWithPagination(db, page, size)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"total": total,
		"page":  page,
		"size":  size,
		"data":  products,
	})
}
