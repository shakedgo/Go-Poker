package main

import (
	"fmt"
	"log"

	"github.com/shakedgo/Go-Poker/Backend/config"
	"github.com/shakedgo/Go-Poker/Backend/pkg/api"
	"github.com/shakedgo/Go-Poker/Backend/pkg/db"
)

func main() {
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
