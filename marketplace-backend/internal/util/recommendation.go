package util

import (
	"gorm.io/gorm"
	"marketplace-backend/internal/model"
)

// GenerateRecommendations dynamically fetches listings based on user activity and interests.
func GenerateRecommendations(db *gorm.DB, userID uint) ([]model.Listing, error) {
	var recommendedListings []model.Listing
	var userInterests []model.UserInterest // Assuming a model.UserInterest that tracks user interests or interactions

	// Fetch user interests or tags from their viewed or favorited listings
	if err := db.Model(&model.UserInterest{}).Where("user_id = ?", userID).Find(&userInterests).Error; err != nil {
		return nil, err
	}

	// Convert user interests into a slice of IDs or tags for querying
	interestIDs := make([]uint, 0, len(userInterests))
	for _, interest := range userInterests {
		interestIDs = append(interestIDs, interest.TagID) // Assuming TagID is an attribute of UserInterest
	}

	if len(interestIDs) == 0 {
		// Fallback to a default recommendation logic if no specific interests are found
		if err := db.Limit(10).Find(&recommendedListings).Error; err != nil {
			return nil, err
		}
		return recommendedListings, nil
	}

	// Example logic: Fetch listings associated with the user's interests or tags
	err := db.Debug().Model(&model.Listing{}).
		Joins("JOIN listing_tags ON listing_tags.listing_id = listings.id").
		Where("listing_tags.tag_id IN ?", interestIDs).
		Distinct().
		Limit(10). // Limit the number of recommendations for performance reasons
		Find(&recommendedListings).Error

	if err != nil {
		return nil, err
	}

	// This example can be further refined with additional filters, such as excluding listings the user has already interacted with.

	return recommendedListings, nil
}
