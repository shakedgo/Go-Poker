package api

import (
	"Go-Poker/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//	func addPlayer(c *gin.Context) {
//		var playerName string
//		fmt.Println("Type a name")
//		fmt.Scanln(&playerName)
//		_ = utils.NewPlayer(playerName)
//	}
type AddPlayerRequest struct {
	Name string `json:"name"`
}

func AddPlayer(c *gin.Context) {
	// Parse JSON request body
	var req AddPlayerRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validate player name
	if req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Player name cannot be empty"})
		return
	}

	// Create a new player with the provided name
	player := utils.NewPlayer(req.Name)
	// if player == nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create player"})
	// 	return
	// }

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "Player added successfully", "player": player})
}

func JoinTable(c *gin.Context) {
	tableID, err := strconv.Atoi(c.Param("tableid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}
	table := utils.GetTableByID(tableID)
	if table == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve table"})
		return
	}
	playerID, err := strconv.Atoi(c.Param("playerid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}
	player := utils.GetPlayerByID(playerID)
	if player == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve player"})
		return
	}
	table.JoinTable(player)
}

func PrintTable(c *gin.Context) {
	tableID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}
	table := utils.GetTableByID(tableID)
	if table == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve table"})
		return
	}

	c.JSON(http.StatusOK, table.String())
	// fmt.Println(table)
}
