package controllers

import (
	"internTest01/pkg/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserController(routerGroup *gin.RouterGroup, db *gorm.DB) {
	userRoute := routerGroup.Group("/user")
	userRoute.GET("/:id", func(context *gin.Context) {
		services.GetUserById(context, db)
	})
}
