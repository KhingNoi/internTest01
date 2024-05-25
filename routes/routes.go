package routes

import (
	"internTest01/pkg/controllers"

	"github.com/gin-gonic/gin"
)

type AppRouter struct {
	route *gin.Engine
}

func Router() AppRouter {
	router := AppRouter{route: gin.Default()}
	main := router.route.Group("/")
	// api := router.route.Group("/api")

	controllers.HelpCheckController(main)
	return router
}

func (router AppRouter) Run(addr ...string) error {
	return router.route.Run()
}
