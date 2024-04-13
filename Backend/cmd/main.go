package main

import (
	"Go-Poker/pkg/api"
	"Go-Poker/pkg/db"
	"Go-Poker/pkg/utils"
	"context"
	"log"
)

func main() {
	db.Connect()
	router := api.SetupRouter()
	utils.InitTable()
	// non blocking operation
	// go func() {
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start the server:", err)
	}
	// }()

	// Close the database connection when the application exits
	defer func() {
		if err := db.Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// // Wait 3 seconds for the server to start
	// time.Sleep(3 * time.Second)
}
