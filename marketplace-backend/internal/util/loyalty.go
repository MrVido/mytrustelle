package util

import (
	"marketplace-backend/internal/model"
	"gorm.io/gorm"
)

// AwardPointsToUser adds loyalty points to a user's account after a transaction.
func AwardPointsToUser(db *gorm.DB, userID uint, pointsAwarded int) error {
	// Example: Update the user's loyalty points balance
	// This function assumes the existence of a User model with a Points field

	var user model.User
	if err := db.First(&user, userID).Error; err != nil {
		return err
	}

	user.Points += pointsAwarded
	return db.Save(&user).Error
}
