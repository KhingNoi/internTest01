package repositories

import (
	"internTest01/pkg/models"

	"gorm.io/gorm"
)

func FindProductById(db *gorm.DB, id string) (models.Product, error) {
	var product models.Product
	result := db.First(&product, id)
	return product, result.Error
}

func FindProductListWithPagination(db *gorm.DB, page int, size int) ([]models.Product, int64, error) {
	var productsWithPageAndSize []models.Product
	var total int64
	if err := db.Model(&models.Product{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * size
	result := db.Limit(size).Offset(offset).Find(&productsWithPageAndSize)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return productsWithPageAndSize, total, nil
}
