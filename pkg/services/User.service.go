package services

import (
	"internTest01/pkg/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUserById(context *gin.Context, db *gorm.DB) {
	id := context.Param("id")
	if _, strconvErr := strconv.Atoi(id); strconvErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	user, err := repositories.FindUserById(db, id)
	if err != nil {
		if err.Error() == "record not found" {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
	context.JSON(http.StatusOK, user)
}
