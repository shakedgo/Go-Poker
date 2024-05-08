package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartRouter() {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
	})

	router.POST("/login", Login)
	router.POST("/signup", Signup)
	router.POST("/new-player", AddPlayer)
	router.POST("/join-table", JoinTable)
	router.GET("/print-table/:id", PrintTable)
	router.GET("/print-player/:id", PrintPlayer)

	go func() {
		// service connections
		if err := Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

var Server = &http.Server{
	Addr:    ":8080",
	Handler: gin.Default(),
}
