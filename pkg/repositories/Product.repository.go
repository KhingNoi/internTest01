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

func FindProductListWithPagination(db *gorm.DB, page int, size int) ([]models.Product, int, error) {
	var productsWithPageAndSize []models.Product

	offset := (page - 1) * size
	queryResult := db.Limit(size).Offset(offset).Find(&productsWithPageAndSize)
	if queryResult.Error != nil {
		return nil, 0, queryResult.Error
	}
	total := len(productsWithPageAndSize)
	return productsWithPageAndSize, total, nil
}

func FindProductListByName(db *gorm.DB, searchTerm models.SearchTermPayload, page int, size int) ([]models.Product, int, error) {
	var productsByName []models.Product
	offset := (page - 1) * size
	queryResult := db.Where("name LIKE ?", searchTerm.Name+"%").Limit(size).Offset(offset).Find(&productsByName)
	if queryResult.Error != nil {
		return nil, 0, queryResult.Error
	}
	total := len(productsByName)
	return productsByName, total, nil
}
