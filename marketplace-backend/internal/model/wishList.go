package model

import (
    "gorm.io/gorm"
    "time"
)

// WishList represents a collection of items or services that a user has expressed interest in.
type WishList struct {
    gorm.Model
    UserID    uint      `gorm:"not null;index"` // The owner of the wish list
    Name      string    `gorm:"not null"` // Name of the wish list, e.g., "Summer 2024"
    Public    bool      `gorm:"default:false"` // Whether the wish list is visible to other users
    Items     []Listing `gorm:"many2many:wishlist_listings;"` // Associated listings
    CreatedAt time.Time
    UpdatedAt time.Time
}

// WishListListing represents the many-to-many relationship between WishLists and Listings.
// This intermediary table is useful if additional attributes on the relationship are needed.
type WishListListing struct {
    WishListID uint `gorm:"primaryKey"`
    ListingID  uint `gorm:"primaryKey"`
    AddedAt    time.Time // When the item was added to the wish list
}

// Key Enhancements:
// - **Public Visibility**: The `Public` field allows users to control the visibility of their wish lists, supporting both private collections and public curation.
// - **Flexible Association with Listings**: Through a many-to-many relationship, users can easily add or remove listings from their wish lists without affecting the listings themselves.
// - **Date Tracking on Many-to-Many Relationships**: By using an intermediary table (`WishListListing`), the platform can track when items were added, offering insights into user interest over time.

// This model setup not only caters to the fundamental use case of saving favorites but also expands into a more social or shared experience by allowing public wish lists. Additionally, the detailed tracking facilitates a richer data set for understanding user preferences and behaviors.
