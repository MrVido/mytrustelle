package model

import (
    "gorm.io/gorm"
    "time"
)

// ProductReview encapsulates user feedback on products listed on the marketplace.
type ProductReview struct {
    gorm.Model
    ListingID uint      `gorm:"index;not null"` // Associated listing
    UserID    uint      `gorm:"index;not null"` // User who left the review
    Rating    int       `gorm:"type:int;not null"` // Rating given, typically on a scale of 1-5
    Title     string    `gorm:"type:varchar(255)"` // Short summary of the review
    Comment   string    `gorm:"type:text"` // Detailed comment
    Helpful   int       `gorm:"default:0"` // Number of users who found this review helpful
    Reported  bool      `gorm:"default:false"` // Whether the review has been reported as inappropriate
    CreatedAt time.Time
    UpdatedAt time.Time
}

// Key Enhancements:
// - **Title and Comment**: Allows users to provide both a summary and detailed feedback, offering more insight into their rating.
// - **Helpful Votes**: Incorporating a "Helpful" count enables other users to vote on the usefulness of a review, helping to surface the most valuable insights.
// - **Reported Flag**: Adding a flag for reviews that have been reported as inappropriate assists in content moderation efforts.

// This comprehensive model facilitates a rich user review experience, encouraging community engagement and providing valuable insights to potential buyers. It's structured to support not just the collection of feedback, but also its moderation and utility in helping users make informed decisions.
