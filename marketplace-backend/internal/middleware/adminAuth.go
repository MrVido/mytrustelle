package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// AdminRequired is a middleware that checks if the user is an admin.
func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Assumes user role is set in the context from a previous authentication middleware
		// Consider using a more structured approach for context keys to avoid plain string errors
		const userRoleKey = "userRole"
		userRole, exists := c.Get(userRoleKey)
		if !exists {
			// If userRole does not exist in context, it indicates that the authentication middleware did not run or the user is not authenticated
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not authenticated"})
			c.Abort()
			return
		}
		
		// Check if the userRole is indeed "admin"
		if role, ok := userRole.(string); !ok || role != "admin" {
			// If the userRole is not a string or not equal to "admin", deny access
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied. Requires admin role."})
			c.Abort()
			return
		}
		
		// If user is admin, proceed with the request
		c.Next()
	}
}
