package repositories

import (
	"internTest01/pkg/models"

	"gorm.io/gorm"
)

func FindProductById(db *gorm.DB, id string) (models.Product, error) {
	var product models.Product
	queryResult := db.First(&product, id)
	return product, queryResult.Error
}

func FindProductListWithPagination(db *gorm.DB, page int, size int) ([]models.Product, int64, error) {
	var productsWithPageAndSize []models.Product
	var total int64
	if err := db.Model(&models.Product{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * size
	queryResult := db.Limit(size).Offset(offset).Find(&productsWithPageAndSize)
	if queryResult.Error != nil {
		return nil, 0, queryResult.Error
	}
	return productsWithPageAndSize, total, nil
}

// func FindProductByName(db *gorm.DB, name string) ([]models.Product, int64, error) {
// 	var
// }
