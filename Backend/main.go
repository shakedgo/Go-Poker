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
	api.StartRouter()
	utils.InitTable()

	// Close the database connection when the application exits
	defer func() {
		if err := db.Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
		log.Println("Disconnected Database")
	}()
}
