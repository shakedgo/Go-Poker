package main

import (
	"Go-Poker/pkg/api"
	"Go-Poker/pkg/db"
	"Go-Poker/pkg/utils"
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	db.Connect()
	api.StartRouter()
	utils.InitTable()

	// Use a WaitGroup to coordinate between the database and server shutdown
	var wg sync.WaitGroup
	wg.Add(2)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	shutdownComplete := make(chan struct{})

	// Handle shutdown for database
	go func() {
		defer wg.Done()

		<-stopChan // Wait for SIGINT or another signal

		log.Println("Shutdown Database ...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := db.Client.Disconnect(ctx); err != nil {
			log.Fatal("Database Shutdown:", err)
		}
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Timeout of 5 seconds.")
		}
		log.Println("Disconnected Database")

		log.Println("Shutdown HTTP Server ...")
		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := api.Server.Shutdown(ctx); err != nil {
			log.Fatal("HTTP Server Shutdown:", err)
		}
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Timeout of 5 seconds.")
		}
		log.Println("HTTP Server shutdown complete.")
		close(shutdownComplete)
	}()

	// Wait for both shutdowns to complete
	go func() {
		wg.Wait()
		log.Println("All shutdowns complete.")
	}()

	// Wait for the shutdown to complete
	<-shutdownComplete
	log.Println("All shutdowns complete.")
}
