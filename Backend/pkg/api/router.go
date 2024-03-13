package api

import (
	"Go-Poker/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler defines the API request handlers.
type Handler struct {
	// You can include database connections, services, or other dependencies here
	DB *db.Database
}

// NewHandler creates a new instance of the API handler.
func NewHandler(db *db.Database) *Handler {
	return &Handler{
		DB: db,
		// Initialize other dependencies here
	}
}

// PingHandler handles the "GET /ping" endpoint.
func (h *Handler) PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// ExampleHandler handles the "GET /example" endpoint.
func (h *Handler) ExampleHandler(c *gin.Context) {
	// Perform any necessary logic here
	c.JSON(http.StatusOK, gin.H{
		"message": "This is an example response",
	})
}
