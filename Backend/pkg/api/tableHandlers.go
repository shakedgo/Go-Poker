package api

import (
	"Go-Poker/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JoinTableRequest struct {
	TableID  int `json:"TableID"`
	PlayerID int `json:"PlayerID"`
}

func JoinTable(c *gin.Context) {
	var req JoinTableRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if req.TableID == 0 || req.PlayerID == 0 {
		if req.PlayerID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Player id cannot be 0"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Table id cannot be 0"})
		}
		return
	}
	table := utils.GetTableByID(req.TableID)
	if table == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve table"})
		return
	}
	player := utils.GetPlayerByID(req.PlayerID)
	if player == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve player"})
		return
	}
	table.JoinTable(player)

	c.JSON(http.StatusOK, gin.H{"message": "Player Joined Table successfully", "table_id": table.ID, "player": player})

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
