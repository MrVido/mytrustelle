package middleware

import (
	"github.com/gin-gonic/gin"
)

// SecureHeaders adds various security-related headers to the response.
func SecureHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Frame-Options", "DENY") // Prevents clickjacking
		c.Header("X-Content-Type-Options", "nosniff") // Prevents MIME type sniffing
		c.Header("X-XSS-Protection", "1; mode=block") // Enables XSS protection in older browsers
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin") // Controls referrer header information
		c.Header("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline'; object-src 'none';") // Helps prevent XSS attacks
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains") // Enforces HTTPS

		// Proceed to the next middleware/handler
		c.Next()
	}
}
