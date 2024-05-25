package repositories

import (
	"internTest01/pkg/models"

	"gorm.io/gorm"
)

func FindProductList(db *gorm.DB) ([]models.Product, error) {
	var products []models.Product
	result := db.Find(&products)
	return products, result.Error
}
