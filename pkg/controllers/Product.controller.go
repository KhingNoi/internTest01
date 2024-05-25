package controllers

import (
	"internTest01/pkg/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductController(routerGroup *gin.RouterGroup, db *gorm.DB) {
	productRoute := routerGroup.Group("/product")
	productRoute.GET("", func(context *gin.Context) {
		services.GetProductList(context, db)
	})
}