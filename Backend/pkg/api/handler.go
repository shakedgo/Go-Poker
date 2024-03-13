package api

import "github.com/gin-gonic/gin"

func SetupRouter(handler *Handler) *gin.Engine {
	router := gin.Default()

	// Ping endpoint
	router.GET("/ping", handler.PingHandler)

	// Example endpoint
	router.GET("/example", handler.ExampleHandler)

	// Add more endpoints as needed

	return router
}
