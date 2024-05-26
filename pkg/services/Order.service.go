package services

import (
	"internTest01/pkg/models"
	"internTest01/pkg/repositories"
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
		currentStock, err := repositories.FindStockById(db, strconv.Itoa(detail.ProductId))
		if currentStock.Name == "" {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if currentStock.Stock == 0 {
			context.JSON(http.StatusConflict, gin.H{"message": currentStock.Name + " " + "out of stock"})
			return
		}
		if detail.Quantity > currentStock.Stock {
			context.JSON(http.StatusConflict, gin.H{"message": currentStock.Name + " " + "insufficient stock"})
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
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"message": "Order with details created successfully",
	}
	context.JSON(http.StatusOK, response)
}

func UpdateProductStockByAmountRequest(context *gin.Context, db *gorm.DB) {
	id := context.Param("id")
	if _, strconvErr := strconv.Atoi(id); strconvErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	currentOrder, err := repositories.FindOrderById(db, id)
	if err != nil {
		if err.Error() == "record not found" {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if currentOrder.Status != models.PENDING_STATUS {
		context.JSON(http.StatusForbidden, gin.H{"error": "current order status is not Pending"})
		return
	}

	for _, detail := range currentOrder.OrderDetailList {
		product, err := repositories.UpdateProductStock(db, detail.Quantity, detail.ProductId)
		if err != nil {
			if err.Error() == "out of stock" || err.Error() == "insufficient stock" {
				context.JSON(http.StatusConflict, gin.H{"message": product.Name + " " + err.Error()})
				return
			} else {
				context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
	}

	updateErr := repositories.UpdateOrderStatus(db, currentOrder)
	if updateErr != nil {
		if updateErr.Error() == "record not found" {
			context.JSON(http.StatusNotFound, gin.H{"error": updateErr.Error()})
			return
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": updateErr.Error()})
			return
		}
	}
	context.JSON(http.StatusOK, gin.H{"message": "order success"})

}
