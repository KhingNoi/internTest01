package database

import (
	"fmt"
	"internTest01/configs"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBInstance *gorm.DB

func ConnectDatabase() *gorm.DB {
	configs.InitEnvConfigs()

	dbConfig := configs.EnvConfigs

	databaseURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.DatabaseUser,
		dbConfig.DatabasePassword,
		dbConfig.DatabaseUrl,
		dbConfig.DatabasePort,
		dbConfig.DatabaseName)

	db, err := gorm.Open(mysql.Open(databaseURI), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("Successfully connected to MySQL database!")

	DBInstance = db
	return db
}

func CloseDB() {
	if DBInstance != nil {
		db, _ := DBInstance.DB()
		db.Close()
		log.Println("Database connection closed.")
	}
}
