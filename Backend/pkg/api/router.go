package api

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/new-player", AddPlayer)
	router.POST("/join-table", JoinTable)
	router.GET("/print-table/:id", PrintTable)
	router.GET("/print-player/:id", PrintPlayer)

	return router
}
