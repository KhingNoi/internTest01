package repositories

import (
	"internTest01/pkg/models"

	"gorm.io/gorm"
)

func FindOrderById(db gorm.DB, id string) (models.Order, error) {
	var order models.Order
	queryResult := db.First(&order, id)
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
