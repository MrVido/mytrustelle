package v1

import (
    "marketplace-backend/internal/model"
    "marketplace-backend/internal/util"
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// LoginUser validates user login credentials and generates JWT tokens upon success.
func LoginUser(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var loginDetails struct {
            Email    string `json:"email"`
            Password string `json:"password"`
        }

        // Bind the JSON to the struct and validate
        if err := c.ShouldBindJSON(&loginDetails); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login details"})
            return
        }

        // Find the user by email
        var user model.User
        result := db.Where("email = ?", loginDetails.Email).First(&user)
        if result.Error != nil || user.Status != "active" { // Check if the user is active
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid login credentials or account inactive"})
            return
        }

        // Check if the password hash matches
        if !util.CheckPasswordHash(loginDetails.Password, user.PasswordHash) {
            // Log the failed login attempt here
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid login credentials"})
            return
        }

        // Generate JWT access and refresh tokens for the authenticated user
        accessToken, refreshToken, err := util.GenerateTokenPair(user.ID)
        if err != nil {
            // Log the error here
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
            return
        }

        // Respond with the tokens
        c.JSON(http.StatusOK, gin.H{
            "message": "Login successful",
            "accessToken": accessToken,
            "refreshToken": refreshToken,
        })
    }
}
