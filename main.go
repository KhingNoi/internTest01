package main

import (
	"internTest01/database"
	"internTest01/routes"
)

func main() {
	db := database.ConnectDatabase()
	defer database.CloseDB()
	routes.Router(db).Run(":8080")
}
