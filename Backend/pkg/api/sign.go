package api

import (
	"Go-Poker/pkg/db"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

	if credentials.Username == "" || credentials.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or password cannot be empty"})
		return
	}

	users := db.Client.Database(db.Name).Collection("users")
	var userFound Credentials
	err := users.FindOne(context.TODO(), bson.D{{Key: "username", Value: credentials.Username}}).Decode(&userFound)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to find user: %v", credentials.Username)})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(credentials.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
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

	users := db.Client.Database(db.Name).Collection("users")
	usersFound, err := users.CountDocuments(context.TODO(), bson.M{"username": credentials.Username})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in checking username"})
		return
	}
	if usersFound > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Username is taken"})
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	credentials.Password = string(hashedPass)

	insertRes, err := users.InsertOne(context.TODO(), credentials)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to insert user: %v", credentials.Username)})
		return
	}

	fmt.Println(insertRes)
	// insertedID := insertRes.InsertedID

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
