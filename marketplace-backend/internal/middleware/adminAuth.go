package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// AdminRequired is a middleware that checks if the user is an admin.
func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole") // Assumes user role is set in the context from a previous auth middleware
		if !exists || userRole != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}
		c.Next()
	}
}
