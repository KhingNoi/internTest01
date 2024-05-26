package controllers

import (
	"internTest01/pkg/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func OrderController(routerGroup *gin.RouterGroup, db *gorm.DB) {
	userRoute := routerGroup.Group("/order")
	userRoute.POST("/create", func(context *gin.Context) {
		services.CreateOrderWithData(context, db)
	})
	userRoute.PATCH("/submit/:id", func(context *gin.Context) {
		services.UpdateProductStockByAmountRequest(context, db)
	})
}
