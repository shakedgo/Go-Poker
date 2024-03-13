package main

import (
	"Go-Poker/config"
	"Go-Poker/pkg/api"
	"Go-Poker/pkg/db"
	"Go-Poker/pkg/utils"
	"fmt"
	"log"
)

func main() {
	str := utils.ExampleHelperFunction("dsa")
	fmt.Println(str)
	// Load configurations
	appConfig := config.NewConfig()

	// Initialize the database connection
	db, err := db.NewDatabase(
		appConfig.DBHost,
		appConfig.DBPort,
		appConfig.DBUser,
		appConfig.DBPassword,
		appConfig.DBName,
	)
	if err != nil {
		log.Fatal("Failed to initialize the database:", err)
	}

	// Create an instance of the API handler with dependencies
	handler := api.NewHandler(db)

	// Set up the Gin router
	router := api.SetupRouter(handler)

	// Run the server
	addr := fmt.Sprintf(":%s", appConfig.ServerPort)
	err = router.Run(addr)
	if err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
