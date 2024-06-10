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
	if _, err := strconv.Atoi(id); err != nil {
		HandleError(context, http.StatusBadRequest, "Invalid ID format")
		return
	}

	user, err := repositories.FindUserById(db, id)
	if err != nil {
		HandleDatabaseError(context, err)
		return
	}

	context.JSON(http.StatusOK, user)
}
