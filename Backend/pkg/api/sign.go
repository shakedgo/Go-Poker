package api

import (
	"Go-Poker/pkg/db"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var JwtKey = []byte(os.Getenv("JWT_Key"))

type Credentials struct {
	ID       string `bson:"_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
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

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userFound.ID, // Use the user ID from the database
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "jwt_token",
		Value:    tokenString,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func Logout(c *gin.Context) {
	// Get the JWT token from the request cookies
	_, err := c.Request.Cookie("jwt_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No JWT token found"})
		return
	}

	// Clear the JWT token cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "jwt_token",
		Value:    "",
		Expires:  time.Unix(0, 0), // Set the expiration time to the past
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
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
