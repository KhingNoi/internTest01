package main

import (
	"fmt"
	"internTest01/configs"
	"internTest01/routes"
)

func main() {
	configs.InitEnvConfigs()
	testEnv1 := configs.EnvConfigs.DatabaseName
	testEnv2 := configs.EnvConfigs.DatabasePassword
	testEnv3 := configs.EnvConfigs.DatabasePort
	testEnv4 := configs.EnvConfigs.DatabaseUrl
	testEnv5 := configs.EnvConfigs.DatabaseUser
	fmt.Println("DatabaseName :", testEnv1)
	fmt.Println("DatabasePassword :", testEnv2)
	fmt.Println("DatabasePort :", testEnv3)
	fmt.Println("DatabaseUrl :", testEnv4)
	fmt.Println("DatabaseUser :", testEnv5)
	routes.Router().Run(":8080")
}
