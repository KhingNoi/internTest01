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
		HandleError(context, http.StatusBadRequest, "Invalid JSON")
		return
	}

	var orderDetails []*models.OrderDetails
	for _, detail := range orderRequestData.OrderDetailList {
		currentStock, err := repositories.FindStockById(db, strconv.Itoa(detail.ProductId))
		if err != nil {
			HandleDatabaseError(context, err)
			return
		}
		if currentStock.Name == "" {
			HandleError(context, http.StatusNotFound, "Product not found")
			return
		}
		if currentStock.Stock == 0 {
			HandleError(context, http.StatusConflict, currentStock.Name+" out of stock")
			return
		}
		if detail.Quantity > currentStock.Stock {
			HandleError(context, http.StatusConflict, currentStock.Name+" insufficient stock")
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
		HandleDatabaseError(context, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Order with details created successfully"})
}

func UpdateProductStockByAmountRequest(context *gin.Context, db *gorm.DB) {
	id := context.Param("id")
	if _, err := strconv.Atoi(id); err != nil {
		HandleError(context, http.StatusBadRequest, "Invalid ID format")
		return
	}

	currentOrder, err := repositories.FindOrderById(db, id)
	if err != nil {
		HandleDatabaseError(context, err)
		return
	}

	if currentOrder.Status != models.PENDING_STATUS {
		HandleError(context, http.StatusForbidden, "Current order status is not Pending")
		return
	}

	for _, detail := range currentOrder.OrderDetailList {
		product, err := repositories.UpdateProductStock(db, detail.Quantity, detail.ProductId)
		if err != nil {
			if err.Error() == "out of stock" || err.Error() == "insufficient stock" {
				HandleError(context, http.StatusConflict, product.Name+" "+err.Error())
				return
			} else {
				HandleDatabaseError(context, err)
				return
			}
		}
	}

	err = repositories.UpdateOrderStatus(db, currentOrder)
	if err != nil {
		HandleDatabaseError(context, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Order success"})
}
