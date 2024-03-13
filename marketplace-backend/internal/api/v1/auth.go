package v1

import (
    "marketplace-backend/internal/model"
    "marketplace-backend/internal/util"
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// LoginUser validates user login credentials and generates a JWT token upon success.
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
        if result.Error != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid login credentials"})
            return
        }

        // Check if the password hash matches
        if !util.CheckPasswordHash(loginDetails.Password, user.PasswordHash) {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid login credentials"})
            return
        }

        // Generate JWT token for the authenticated user
        token, err := util.GenerateToken(user.ID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
            return
        }

        // Respond with the token
        c.JSON(http.StatusOK, gin.H{
            "message": "Login successful",
            "token": token,
        })
    }
}
