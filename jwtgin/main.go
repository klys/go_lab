package main

import (
	//"fmt"
	"net/http"
	//"time"

	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("your-secret-key")

func main() {
	r := gin.Default()

	// Middleware for JWT token authentication
	r.Use(authMiddleware)

	r.GET("/public", publicEndpoint)
	r.GET("/private", privateEndpoint)

	r.Run(":8080")
}

func authMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Next()
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	c.Abort()
}

func publicEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This is a public endpoint"})
}

func privateEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This is a private endpoint"})
}
