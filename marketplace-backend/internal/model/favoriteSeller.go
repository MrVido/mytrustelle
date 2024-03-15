package model

import (
	"gorm.io/gorm"
	"time"
)

// FavoriteSeller represents the relationship between a user and their favorite sellers on the marketplace.
type FavoriteSeller struct {
	gorm.Model
	UserID         uint      `gorm:"not null;index"`           // ID of the user who has favorited a seller
	SellerID       uint      `gorm:"not null;index"`           // ID of the seller being favorited
	FavoritedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"` // Timestamp of when the seller was favorited
	NotifyNewItems bool      `gorm:"default:false"`            // Whether the user opts in to receive notifications for new items from the seller
}

// Hooks to automatically set FavoritedAt
func (fs *FavoriteSeller) BeforeCreate(tx *gorm.DB) (err error) {
	fs.FavoritedAt = time.Now()
	return
}

// You might consider implementing methods to enable or disable notifications for new items.
// This can involve updating the NotifyNewItems field based on user preferences.

// Enhancements:

// FavoritedAt Timestamp: The addition of the FavoritedAt field allows the platform to track when a user favorites a seller, providing insights into user engagement and potential seasonal preferences.

// Notification Preferences: The NotifyNewItems boolean field enables users to opt-in for notifications about new listings from their favorite sellers. This feature enhances the user experience by keeping them informed about fresh items that match their interests.

// Automatic Timestamping: The BeforeCreate hook automatically sets the FavoritedAt timestamp to the current time when a new favorite seller relationship is created, ensuring accurate tracking of when the relationship was established.

// Dynamic Engagement: With the NotifyNewItems field, the platform can send targeted notifications or emails to users about new listings from favorite sellers, driving engagement and potentially increasing sales.

// This enhanced FavoriteSeller model not only allows users to keep track of their preferred sellers but also opens avenues for increased platform engagement through targeted notifications, creating a more personalized and dynamic shopping experience.





