package middleware

import (
	"marketplace-backend/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JWTAuthentication validates JWT token in the request header.
func JWTAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			c.Abort()
			return
		}

		claims, err := util.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Pass the user ID to the next middleware or handler
		c.Set("userID", claims.UserID)

		c.Next()
	}
}
