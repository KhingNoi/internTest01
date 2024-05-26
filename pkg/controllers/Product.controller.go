package controllers

import (
	"internTest01/pkg/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductController(routerGroup *gin.RouterGroup, db *gorm.DB) {
	productRoute := routerGroup.Group("/product")

	productRoute.GET("/", func(context *gin.Context) {
		services.GetProductListWithPagination(context, db)
	})
	productRoute.GET("/:id", func(context *gin.Context) {
		services.GetProductById(context, db)
	})
	productRoute.GET("/search", func(context *gin.Context) {
		services.SearchProductListByName(context, db)
	})
	productRoute.GET("/checkStock/:id", func(context *gin.Context) {
		services.CheckStockByProductId(context, db)
	})

}
