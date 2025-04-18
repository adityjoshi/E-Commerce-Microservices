package middleware

import (
	"net/http"

	"github.com/adityjoshi/E-Commerce-Microservices/service/authService/utils"
	"github.com/gin-gonic/gin"
)

func AuthorizeUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing"})
			c.Abort()
			return
		}

		claims, err := utils.DecodeJwt(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Extract user_type
		userType, ok := claims["User_type"].(string)
		if !ok || userType != "User" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Extract region
		if region, exists := claims["region"].(string); exists {
			c.Set("region", region)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Region not specified in token"})
			c.Abort()
			return
		}

		// Extract ID and ensure it's uint
		var userID uint
		if id, exists := claims["ID"].(float64); exists {
			userID = uint(id) // Convert float64 to uint
			c.Set("ID", userID)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "ID not specified in token"})
			c.Abort()
			return
		}

		// Extract Email
		if email, exists := claims["Email"].(string); exists {
			c.Set("Email", email)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Email not specified in token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
