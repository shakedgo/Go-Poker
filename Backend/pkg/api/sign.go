package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var Users []Credentials

const (
	MinCost     int = 4  // the minimum allowable cost as passed in to GenerateFromPassword
	MaxCost     int = 31 // the maximum allowable cost as passed in to GenerateFromPassword
	DefaultCost int = 10 // the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword
)

func Login(c *gin.Context) {
	var credentials Credentials

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	for _, user := range Users {
		if user.Username == credentials.Username {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
			if err == nil {
				c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
				return
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Password"})
				return
			}
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Username"})
}

func Signup(c *gin.Context) {
	var credentials Credentials

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if credentials.Username == "" || credentials.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or password cannot be empty"})
		return
	}

	for _, user := range Users {
		if user.Username == credentials.Username {
			c.JSON(http.StatusConflict, gin.H{"error": "Username is taken"})
			return
		}
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	credentials.Password = string(hashedPass)
	Users = append(Users, credentials)

	c.JSON(http.StatusCreated, gin.H{"message": "Signup successful"})
}
