package api

import (
	"Go-Poker/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AddPlayerRequest struct {
	Name string `json:"name"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var Users []Credentials

// Users := []Credentials{
// 	{username: "Shaked", password: "12345"},
// 	{username: "Emanu", password: "12345"},
// }
// Users = append(Users, credentials{username: "shaked", password: "12345"})

/////////
// GET //
/////////

func PrintPlayer(c *gin.Context) {
	playerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}
	player := utils.GetPlayerByID(playerID)
	if player == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve player"})
		return
	}

	c.JSON(http.StatusOK, player.String())
	// fmt.Println(table)
}

//////////
// POST //
//////////

func AddPlayer(c *gin.Context) {
	var req AddPlayerRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Player name cannot be empty"})
		return
	}

	player := utils.NewPlayer(req.Name)

	// Success message
	c.JSON(http.StatusOK, gin.H{"message": "Player added successfully", "player": player})
}

func Login(c *gin.Context) {
	var credentials Credentials

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	for _, user := range Users {
		if user.Username == credentials.Username && user.Password == credentials.Password {
			c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

func Signup(c *gin.Context) {
	var credentials Credentials

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	Users = append(Users, credentials)
}
