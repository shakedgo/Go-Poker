package api

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
	})

	router.POST("/login", Login)
	router.POST("/new-player", AddPlayer)
	router.POST("/join-table", JoinTable)
	router.GET("/print-table/:id", PrintTable)
	router.GET("/print-player/:id", PrintPlayer)

	return router
}
