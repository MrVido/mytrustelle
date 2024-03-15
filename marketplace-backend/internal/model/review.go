package model

import (
	"gorm.io/gorm"
)

// Review represents feedback left by a user for another user or a listing.
// Enhancements include moderation flags, response capability, and review helpfulness votes.
type Review struct {
	gorm.Model
	ReviewerID   uint           `gorm:"index;not null"` // Index for faster query performance
	RevieweeID   uint           `gorm:"index;not null"` // Index for aggregating reviews
	ListingID    uint           `gorm:"index;not null"` // Allows reviews to be associated with listings
	Rating       int            `gorm:"not null;check:rating >= 1 AND rating <= 5"` // Enforce rating bounds (1-5)
	Comment      string         `gorm:"type:text"`      // Optional comment with the review
	CreatedAt    time.Time
	IsModerated  bool           `gorm:"default:false"`  // Indicates if the review has been checked by moderators
	Response     *ReviewResponse `gorm:"foreignKey:ReviewID"` // Optional response to the review
	HelpfulVotes int            `gorm:"default:0"`      // Count of users who found the review helpful
	Reported     bool           `gorm:"default:false"`  // Indicates if the review has been reported for moderation
}

// ReviewResponse represents a response from the reviewee to a particular review.
// This allows for a two-way conversation and clarifications if necessary.
type ReviewResponse struct {
	gorm.Model
	ReviewID    uint   `gorm:"uniqueIndex;not null"`
	ResponderID uint   `gorm:"not null"` // Should match RevieweeID in the associated review
	Content     string `gorm:"type:text;not null"`
	CreatedAt   time.Time
}

// Key Enhancements:

// 1. **Rating Validation**: Ensure ratings are within a specific range (1-5) to maintain consistency.

// 2. **Moderation and Reporting**: Flags like `IsModerated` and `Reported` help manage content quality and community standards.

// 3. **Review Responses**: Allowing reviewees to respond to reviews fosters dialogue and can provide valuable context to other users.

// 4. **Helpfulness Votes**: The `HelpfulVotes` count enables users to signal the usefulness of a review, improving the overall quality of feedback.

// 5. **Database Indexes**: Adding indexes on `ReviewerID`, `RevieweeID`, and `ListingID` optimizes queries related to reviews, especially for aggregation and filtering purposes.

// These enhancements ensure the review system is not only a platform for feedback but also an interactive space that supports moderation, engagement, and valuable insights for users making decisions based on reviews.
