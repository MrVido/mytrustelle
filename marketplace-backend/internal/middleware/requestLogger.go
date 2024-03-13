package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestLogger logs details about each HTTP request
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Log request details after response has been sent
		duration := time.Since(start)
		statusCode := c.Writer.Status()
		log.Printf("[%s] %s %s %d %s", c.Request.Method, c.Request.RequestURI, c.ClientIP(), statusCode, duration)
	}
}
