package model

import (
	"gorm.io/gorm"
)

// Base Listing model to represent the common attributes of all listings.
type Listing struct {
	gorm.Model
	UserID      uint      `gorm:"not null"` // The user who posted the listing
	Title       string    `gorm:"not null"` // Listing title
	Description string    `gorm:"type:text;not null"` // Description of the listing
	Price       float64   `gorm:"not null"` // Price of the item or service
	Category    Category  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // The category of the listing
	Tags        []Tag     `gorm:"many2many:listing_tags;"` // Tags for search optimization and categorization
	// Polymorphic association fields
	ListableID   uint   `gorm:"not null"`
	ListableType string `gorm:"not null"`
}

// Category model to categorize listings
type Category struct {
	gorm.Model
	Name        string    `gorm:"unique;not null"` // Unique name for the category
	Description string    `gorm:"type:text;not null"` // Description of the category
	Listings    []Listing `gorm:"foreignKey:CategoryID"`
}

// Tag model for tagging listings for better searchability
type Tag struct {
	gorm.Model
	Name     string    `gorm:"unique;not null"` // Tag name
	Listings []Listing `gorm:"many2many:listing_tags;"`
}

// Examples of specific listing types
// VehicleListing for listings in the "Vehicles" category
type VehicleListing struct {
	gorm.Model
	ListingID uint // Associated Listing ID
	Make      string
	Model     string
	Year      int
	Mileage   int
}

// PropertyListing for listings in the "Real Estate" category
type PropertyListing struct {
	gorm.Model
	ListingID   uint // Associated Listing ID
	PropertyType string
	SquareFeet   int
	Bedrooms     int
	Bathrooms    int
}

// ServiceListing for listings in the "Services" category
type ServiceListing struct {
	gorm.Model
	ListingID uint // Associated Listing ID
	ServiceType string
	Hours      float64
}

// ElectronicListing for listings in the "Electronics" category
type ElectronicListing struct {
	gorm.Model
	ListingID uint // Associated Listing ID
	DeviceType string
	Brand      string
	Model      string
	Condition  string
}

// Implement polymorphic association in GORM by defining ListableID and ListableType in the Listing model, allowing a listing to be associated with any type of specific listing detail (e.g., VehicleListing, PropertyListing).

// This setup allows the core Listing model to remain generic and applicable across a wide range of categories, while specific details related to different categories can be encapsulated in their respective models.

// Additionally, consider implementing interfaces or service layers in your application logic to handle the specific business rules and validations required for each listing type, ensuring a flexible yet structured approach to managing diverse listings in your marketplace.
