package model

import (
	"gorm.io/gorm"
)

// User represents a user in the marketplace with detailed attributes.
type User struct {
	gorm.Model
	Username            string    `gorm:"uniqueIndex;not null"`
	Email               string    `gorm:"uniqueIndex;not null"`
	PhoneNumber         string    `gorm:"index"` // Removed not null constraint to allow flexibility
	PasswordHash        string    `gorm:"not null"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	Role                string    `gorm:"default:user"` // "user", "admin", or other roles
	EmailPreferences    EmailPreferences `gorm:"embedded"`
	SocialMediaProfiles SocialMediaProfiles `gorm:"embedded"` // Embed social media profiles
	Location            Location `gorm:"embedded"` // Embed location data for enhanced user profiles
	Bio                 string   `gorm:"type:text"` // A short biography or introduction
	ProfilePictureURL   string   // URL to the user's profile picture
	IsEmailVerified     bool     `gorm:"default:false"` // Track email verification status
}

// EmailPreferences struct defines user preferences for receiving various types of emails.
type EmailPreferences struct {
	Marketing bool `gorm:"default:true"` // Opt-in to marketing emails by default
	Updates   bool `gorm:"default:true"` // Opt-in to platform updates emails by default
}

// SocialMediaProfiles struct encapsulates links to a user's social media profiles.
type SocialMediaProfiles struct {
	Facebook  string
	Twitter   string
	LinkedIn  string
	Instagram string
	YouTube   string
}

// Location struct stores structured user location data.
type Location struct {
	Country     string
	State       string
	City        string
	PostalCode  string
	AddressLine string
}

// Key Enhancements:

// 1. **Flexible Contact Information**: Making the phone number optional allows users to choose their preferred method of contact.

// 2. **Social Media Integration**: `SocialMediaProfiles` encourage users to connect their accounts, fostering community engagement and verification.

// 3. **Location Information**: Incorporating structured location data (`Location`) can facilitate local search, regional trends, and compliance with location-based regulations.

// 4. **Bio and Profile Picture**: Adding a bio and profile picture URL personalizes user profiles, improving the social aspects of the marketplace.

// 5. **Email Verification**: The `IsEmailVerified` flag helps enhance security and trustworthiness within the platform by encouraging verified accounts.

// 6. **Embedded Preferences and Profiles**: Using GORM's embedded structs for email preferences, social profiles, and location keeps the model organized and extensible.

// These enhancements aim to build a more engaging, secure, and user-friendly platform that respects user preferences and encourages community interaction.
