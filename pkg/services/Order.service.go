package services

import (
	"internTest01/pkg/models"
	"internTest01/pkg/repositories"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrderWithData(context *gin.Context, db *gorm.DB) {
	var orderRequestData models.OrderPayload
	if err := context.BindJSON(&orderRequestData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	var orderDetails []*models.OrderDetails
	for _, detail := range orderRequestData.OrderDetailList {
		currentStock, _ := repositories.FindStockById(db, strconv.Itoa(detail.ProductId))
		if currentStock.Stock == 0 {
			context.JSON(http.StatusConflict, gin.H{"message": currentStock.Name + " out of stock"})
			return
		}
		if detail.Quantity > currentStock.Stock {
			context.JSON(http.StatusConflict, gin.H{"message": "amount of " + currentStock.Name + " more than currently in stock."})
			return
		}
		orderDetails = append(orderDetails, &models.OrderDetails{
			ProductId: detail.ProductId,
			Quantity:  detail.Quantity,
			Price:     detail.Price,
		})
	}

	order := &models.Order{
		UserID:    orderRequestData.UserId,
		OrderDate: time.Now(),
		Status:    models.PENDING_STATUS,
		Total:     0,
	}

	var totalResult float64
	for _, detail := range orderDetails {
		totalResult += detail.Price
	}
	order.Total = totalResult

	err := repositories.CreateOrderWithDetails(db, order, orderDetails)
	if err != nil {
		log.Fatal("failed to create order with details:", err)
	}

	response := gin.H{
		"message": "Order with details created successfully",
		"order":   order,
		"details": orderDetails,
	}
	context.JSON(http.StatusOK, response)
}
