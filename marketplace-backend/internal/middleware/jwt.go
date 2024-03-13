package middleware

import (
	"strings"
	"marketplace-backend/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JWTAuthentication validates JWT token in the request header.
func JWTAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Expect token to be in the format `Bearer <TOKEN>`
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer <token>"})
			c.Abort()
			return
		}

		tokenString := bearerToken[1]
		claims, err := util.ValidateToken(tokenString)
		if err != nil {
			// Provide more detailed error messages based on the validation issue
			var errMsg string
			if err == util.ErrTokenExpired {
				errMsg = "Token has expired"
			} else if err == util.ErrTokenInvalid {
				errMsg = "Token is invalid"
			} else {
				errMsg = "Error validating token"
			}

			c.JSON(http.StatusUnauthorized, gin.H{"error": errMsg})
			c.Abort()
			return
		}

		// Pass the claims to the next middleware or handler, which could include user roles, permissions, etc.
		c.Set("userID", claims.UserID)
		// If your claims include roles or permissions, you could set them here as well:
		// c.Set("roles", claims.Roles)

		c.Next()
	}
}
