package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header missing"})
		c.Abort()
		return
	}

	// Validate the input token
	if !validateInputToken(tokenString) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid authorization header format"})
		c.Abort()
		return
	}

	// Parse the token, and extract the claims
	// codes..

	// Call the next handler
	c.Next()
}

func validateInputToken(s string) bool {
	// Check if the input string is empty
	if s == "" {
		return false
	}

	fields := strings.Fields(s)
	// Check if the input string has exactly 2 fields
	if len(fields) != 2 {
		return false
	}
	// Check if the first field is "Bearer"
	if fields[0] != "Bearer" {
		return false
	}

	// check for 2nd part
	parts := strings.Split(fields[1], ".")
	// Check if there are exactly 3 parts
	if len(parts) != 3 {
		return false
	}
	// Check if the middle part is not an empty string
	if parts[1] == "" {
		return false
	}

	// Return true if all conditions are satisfied
	return true
}
