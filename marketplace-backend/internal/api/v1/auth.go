package v1

import (
    "marketplace-backend/internal/model"
    "marketplace-backend/internal/util"
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// LoginUser validates user login credentials.
func LoginUser(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var loginDetails struct {
            Email    string `json:"email"`
            Password string `json:"password"`
        }
        if err := c.ShouldBindJSON(&loginDetails); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login details"})
            return
        }

        var user model.User
        result := db.Where("email = ?", loginDetails.Email).First(&user)
        if result.Error != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid login credentials"})
            return
        }

        if !util.CheckPasswordHash(loginDetails.Password, user.PasswordHash) {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid login credentials"})
            return
        }
		// In your LoginUser function, after successfully validating the user's password:
		token, err := util.GenerateToken(user.ID)
		if err != nil {
   		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
    	return
}
		c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})


        // Placeholder for token generation
        c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
    }
}
