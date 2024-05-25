package routes

import (
	"internTest01/pkg/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppRouter struct {
	route *gin.Engine
}

func Router(db *gorm.DB) AppRouter {
	router := AppRouter{route: gin.Default()}
	main := router.route.Group("/")
	api := router.route.Group("/api")

	controllers.HelpCheckController(main)
	controllers.ProductController(api, db)
	return router
}

func (router AppRouter) Run(addr ...string) error {
	return router.route.Run()
}
