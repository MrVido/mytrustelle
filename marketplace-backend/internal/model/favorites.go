package model

import (
	"gorm.io/gorm"
)

// Favorite represents a user's saved favorite listing, enhancing the browsing experience.
type Favorite struct {
	gorm.Model
	UserID    uint `gorm:"index;not null"` // The user who favorited the listing
	ListingID uint `gorm:"index;not null"` // The listing that was favorited
	// Consider adding additional fields for categorization or notes
}

// Key Enhancements:
// - Utilize GORM’s indexing on `UserID` and `ListingID` to improve performance for queries fetching a user’s favorites or listings marked as favorite by multiple users.

// Both enhancements to the `Notifications` and `Favorites` models 
// not only add valuable features to your marketplace platform but also aim to increase user engagement
// and satisfaction by providing personalized interactions and helping users keep track of items they are interested in.