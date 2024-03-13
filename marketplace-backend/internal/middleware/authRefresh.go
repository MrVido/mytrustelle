package middleware

import (
    "net/http"
    "time"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "marketplace-backend/internal/util"
)

func TokenRefreshRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            // If no token is present, proceed without action
            c.Next()
            return
        }

        // Assuming your tokens are formatted as 'Bearer <token>'
        tokenString = tokenString[len("Bearer "):]

        // Parse the token
        token, err := jwt.ParseWithClaims(tokenString, &util.Claims{}, func(token *jwt.Token) (interface{}, error) {
            return util.JwtSecret, nil
        })

        if claims, ok := token.Claims.(*util.Claims); ok && token.Valid {
            // Check how close to expiration the token is
            const refreshThreshold = 30 // minutes
            remaining := time.Until(time.Unix(claims.ExpiresAt, 0))
            if remaining < time.Minute*time.Duration(refreshThreshold) {
                // Token is within the refresh threshold
                // Notify the client that token refresh is recommended
                c.Header("X-Token-Refresh", "Recommended")
            }
        } else if err != nil {
            // Handle parsing errors, but don't interrupt the request
            c.Header("X-Token-Refresh", "Invalid")
        }

        c.Next()
    }
}
