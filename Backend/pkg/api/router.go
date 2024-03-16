package api

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/new-player", AddPlayer)
	router.GET("/join-table/:tableid/:playerid", JoinTable)
	router.GET("/print-table/:id", PrintTable)

	return router
}
