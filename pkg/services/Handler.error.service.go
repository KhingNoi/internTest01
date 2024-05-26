package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(context *gin.Context, statusCode int, message string) {
	context.JSON(statusCode, gin.H{"error": message})
}

func HandleDatabaseError(context *gin.Context, err error) {
	if err.Error() == "record not found" {
		HandleError(context, http.StatusNotFound, "Record not found")
	} else {
		HandleError(context, http.StatusInternalServerError, "Internal server error")
	}
}
