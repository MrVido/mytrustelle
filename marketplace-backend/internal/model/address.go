package model

import (
    "gorm.io/gorm"
)

// Address represents a shipping or billing address for a user on the marketplace platform.
type Address struct {
    gorm.Model
    UserID     uint   `gorm:"index;not null"` // Associate with a user
    FullName   string `gorm:"not null"`       // Full name for shipping
    Line1      string `gorm:"not null"`       // Address line 1
    Line2      string                         // Address line 2 (optional)
    City       string `gorm:"not null"`       // City
    State      string `gorm:"not null"`       // State/Province/Region
    PostalCode string `gorm:"not null"`       // Postal/ZIP code
    Country    string `gorm:"not null"`       // Country
    IsDefault  bool   `gorm:"default:false"`  // Indicates if this is the default address for the user
    AddressType string `gorm:"type:varchar(20);not null"` // "shipping" or "billing"
}

// Key Enhancements:
// - **AddressType**: Distinguishes between shipping and billing addresses, providing flexibility for platforms that require separate addresses for billing and delivery.
// - **IsDefault Field**: Helps users manage multiple addresses by marking one as the default, simplifying the checkout process.
// - **Comprehensive Structure**: Includes all necessary fields to cover a wide range of global addresses, ensuring compatibility and user convenience.

// This Address model provides a solid foundation for handling complex shipping and billing requirements, enhancing the user experience by allowing detailed address management directly within the platform.
