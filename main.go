package main

import (
	"internTest01/routes"
)

func main() {
	routes.Router().Run(":8080")
}
