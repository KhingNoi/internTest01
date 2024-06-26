package services

import (
	"internTest01/pkg/models"
	"internTest01/pkg/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProductById(context *gin.Context, db *gorm.DB) {
	id := context.Param("id")
	if _, err := strconv.Atoi(id); err != nil {
		HandleError(context, http.StatusBadRequest, "Invalid ID format")
		return
	}
	product, err := repositories.FindProductById(db, id)
	if err != nil {
		HandleDatabaseError(context, err)
		return
	}
	context.JSON(http.StatusOK, product)
}

func GetProductListWithPagination(context *gin.Context, db *gorm.DB) {
	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(context.DefaultQuery("size", "12"))
	productList, total, err := repositories.FindProductListWithPagination(db, page, size)
	if err != nil {
		HandleDatabaseError(context, err)
		return
	}

	if len(productList) == 0 {
		HandleError(context, http.StatusNoContent, "")
	}

	context.JSON(http.StatusOK, gin.H{
		"total": total,
		"page":  page,
		"size":  size,
		"data":  productList,
	})
}

func SearchProductListByName(context *gin.Context, db *gorm.DB) {
	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(context.DefaultQuery("size", "12"))
	var name models.SearchTermPayload
	if err := context.BindJSON(&name); err != nil {
		HandleError(context, http.StatusBadRequest, "Invalid JSON")
		return
	}

	productList, total, err := repositories.FindProductListByName(db, name, page, size)

	if err != nil {
		HandleDatabaseError(context, err)
		return
	}
	if len(productList) == 0 {
		HandleError(context, http.StatusNotFound, "Record not found")
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"total": total,
		"page":  page,
		"size":  size,
		"data":  productList,
	})
}

func CheckStockByProductId(context *gin.Context, db *gorm.DB) {
	id := context.Param("id")
	if _, err := strconv.Atoi(id); err != nil {
		HandleError(context, http.StatusBadRequest, "Invalid ID format")
		return
	}
	stock, err := repositories.FindStockById(db, id)
	if err != nil {
		HandleDatabaseError(context, err)
		return
	}
	context.JSON(http.StatusOK, stock)
}
