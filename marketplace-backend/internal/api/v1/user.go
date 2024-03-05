import (
    "marketplace-backend/internal/model"
    "marketplace-backend/internal/util" // Import the util package
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// RegisterUser creates a new user in the database.
func RegisterUser(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var newUser model.User
        if err := c.ShouldBindJSON(&newUser); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Hash the user's password
        hashedPassword, err := util.HashPassword(newUser.PasswordHash)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
            return
        }
        newUser.PasswordHash = hashedPassword

        // Save the user with the hashed password to the database
        if result := db.Create(&newUser); result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "userID": newUser.ID})
    }
}
// GetUserProfile fetches the profile of the logged-in user.
func GetUserProfile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}

		var user model.User
		if err := db.First(&user, userID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}
// UpdateUserProfile allows the user to update their profile information.
func UpdateUserProfile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}

		var updateInfo model.User
		if err := c.ShouldBindJSON(&updateInfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Perform the update operation
		db.Model(&model.User{}).Where("id = ?", userID).Updates(updateInfo)

		c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
	}
}
