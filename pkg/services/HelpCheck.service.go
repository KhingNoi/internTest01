package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelpCheckService(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
