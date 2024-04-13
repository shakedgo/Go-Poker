package main

import (
	"Go-Poker/pkg/api"
	"Go-Poker/pkg/db"
	"Go-Poker/pkg/utils"
	"log"
)

func main() {
	router := api.SetupRouter()
	db.Connect()
	utils.InitTable()
	// non blocking operation
	// go func() {
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start the server:", err)
	}
	// }()

	// // Wait 3 seconds for the server to start
	// time.Sleep(3 * time.Second)
}
