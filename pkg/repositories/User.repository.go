package repositories

import (
	"internTest01/pkg/models"

	"gorm.io/gorm"
)

func FindUserById(db *gorm.DB, id string) (models.User, error) {
	var user models.User
	result := db.First(&user, id)
	return user, result.Error
}
