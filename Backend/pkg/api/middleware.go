package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func authMiddleware(c *gin.Context) {
	cookie, err := c.Request.Cookie("jwt_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token cookie"})
		c.Abort()
		return
	}

	tokenString := cookie.Value

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Set("claims", claims)
	c.Next()
}
