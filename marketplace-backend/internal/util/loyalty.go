package util

import (
	"fmt"
	"marketplace-backend/internal/model"
	"gorm.io/gorm"
)

// AwardPointsToUser adds loyalty points to a user's account after a transaction, with transactional integrity.
func AwardPointsToUser(db *gorm.DB, userID uint, pointsAwarded int) error {
	// Use a transaction to ensure data integrity
	return db.Transaction(func(tx *gorm.DB) error {
		var user model.User
		if err := tx.First(&user, userID).Error; err != nil {
			return fmt.Errorf("user not found: %w", err)
		}

		user.Points += pointsAwarded

		if err := tx.Save(&user).Error; err != nil {
			return fmt.Errorf("failed to update user points: %w", err)
		}

		// Here, you could also insert a record into an audit table or trigger an event for other services,
		// such as sending an email notification about the updated points balance.

		return nil
	})
}
