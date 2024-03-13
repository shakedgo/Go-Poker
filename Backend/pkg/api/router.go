// pkg/api/handler.go

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler defines the API request handlers.
type Handler struct{}

// NewHandler creates a new instance of the API handler.
func NewHandler() *Handler {
	return &Handler{}
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
