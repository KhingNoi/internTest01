package repositories

import (
	"internTest01/pkg/models"

	"gorm.io/gorm"
)

func FindUserById(db *gorm.DB, id string) (models.User, error) {
	var user models.User
	queryResult := db.Preload("Orders").First(&user, id)
	return user, queryResult.Error
}
