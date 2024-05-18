package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Server = &http.Server{
	Addr:    ":8080",
	Handler: gin.Default(),
}

func StartRouter() {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "https://localhost:5173")
		c.Header("Access-Control-Allow-Credentials", "true") // Allow credentials (cookies) in CORS

		// Handle preflight requests (OPTIONS) for all routes
		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Content-Type")
			c.AbortWithStatus(http.StatusOK)
			return
		}

		// Continue processing other requests
		c.Next()
	})

	router.POST("/login", Login)
	router.POST("/signup", Signup)

	protected := router.Group("/")
	protected.Use(authMiddleware)
	{
		router.POST("/logout", Logout)
		protected.POST("/new-player", AddPlayer)
		protected.POST("/join-table", JoinTable)
		protected.GET("/print-table/:id", PrintTable)
		protected.GET("/print-player/:id", PrintPlayer)
	}

	Server.Handler = router

	go func() {
		if err := Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}
