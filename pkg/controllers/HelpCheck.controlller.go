package controllers

import (
	"internTest01/pkg/services"

	"github.com/gin-gonic/gin"
)

func HelpCheckController(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/helpCheck", services.HelpCheckService)
}
