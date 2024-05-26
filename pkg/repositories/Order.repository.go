package repositories

import (
	"errors"
	"internTest01/pkg/models"

	"gorm.io/gorm"
)

func FindOrderById(db *gorm.DB, id string) (models.OrderWithOrderDetail, error) {
	var order models.OrderWithOrderDetail
	queryResult := db.Preload("OrderDetailList").First(&order, id)
	return order, queryResult.Error
}

func CreateOrderWithDetails(db *gorm.DB, order *models.Order, orderDetails []*models.OrderDetails) error {
	transaction := db.Begin()

	if err := transaction.Create(order).Error; err != nil {
		transaction.Rollback()
		return err
	}

	for _, detail := range orderDetails {
		detail.OrderId = order.Id
	}

	if err := transaction.Create(&orderDetails).Error; err != nil {
		transaction.Rollback()
		return err
	}

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		return err
	}

	return nil
}

func UpdateOrderStatus(db *gorm.DB, order models.OrderWithOrderDetail) error {
	order.Status = models.PROCESSING_STATUS
	if err := db.Save(&order).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProductStock(db *gorm.DB, amount int, productId int) (models.Product, error) {
	var product models.Product

	if err := db.First(&product, productId).Error; err != nil {
		return product, err
	}
	if product.Stock == 0 {
		return product, errors.New("out of stock")
	}

	if product.Stock < amount {
		return product, errors.New("insufficient stock")
	}

	product.Stock -= amount
	if err := db.Save(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}
