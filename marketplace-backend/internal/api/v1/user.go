package v1

import (
    "marketplace-backend/internal/model"
    "marketplace-backend/internal/util"
    "net/http"
    "regexp"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// RegisterUser creates a new user in the database with enhanced security and validation.
func RegisterUser(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var newUser model.User
        if err := c.ShouldBindJSON(&newUser); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
            return
        }

        // Email format and password strength validation should be here
        if !IsValidEmail(newUser.Email) {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
            return
        }

        if !IsPasswordStrong(newUser.Password) {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Password does not meet the strength requirements"})
            return
        }

        // Check for existing email
        var existingUser model.User
        if db.Where("email = ?", newUser.Email).First(&existingUser).RowsAffected > 0 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Email already in use"})
            return
        }

        // Hash the user's password
        hashedPassword, err := util.HashPassword(newUser.Password)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
            return
        }
        newUser.PasswordHash = hashedPassword
        newUser.Password = "" // Ensure the plain password is not stored

        // Save the user with the hashed password to the database
        if result := db.Create(&newUser); result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "userID": newUser.ID})
    }
}


// GetUserProfile fetches the profile of the logged-in user with security checks.
func GetUserProfile(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID, exists := c.Get("userID")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
            return
        }

        var user model.User
        if err := db.Select("ID", "Email", "Name").First(&user, userID).Error; err != nil { // Avoid exposing sensitive data
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"user": user})
    }
}

// UpdateUserProfile allows the user to update their profile information securely.
func UpdateUserProfile(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID, exists := c.Get("userID")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
            return
        }

        var updateInfo model.UserUpdate // Assume UserUpdate is a model without sensitive data like PasswordHash
        if err := c.ShouldBindJSON(&updateInfo); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid update data"})
            return
        }

        // Perform the update operation securely, ensuring only certain fields can be updated
        if result := db.Model(&model.User{}).Where("id = ?", userID).Updates(updateInfo); result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
    }
}
func IsValidEmail(email string) bool {
    // This is a simple regex for demonstration. Consider using more comprehensive validation.
    regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
    return regex.MatchString(email)
}

// IsPasswordStrong ensures the password meets defined strength criteria.
func IsPasswordStrong(password string) bool {
    // Example: At least 8 characters, contains a digit and a symbol.
    return len(password) >= 8 && regexp.MustCompile(`[0-9]`).MatchString(password) && regexp.MustCompile(`[!@#$%^&*]`).MatchString(password)
}