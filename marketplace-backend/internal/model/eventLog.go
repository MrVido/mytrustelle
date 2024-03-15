package model

import (
    "gorm.io/gorm"
    "time"
)

// EventType categorizes the event being logged
type EventType string

const (
    EventTypeUserSignup EventType = "user_signup"
    EventTypeUserLogin  EventType = "user_login"
    EventTypeUserUpdate EventType = "user_update"
    EventTypeListingCreated EventType = "listing_created"
    EventTypeListingUpdated EventType = "listing_updated"
    // Add other event types as necessary
)

// EventLog captures significant events in the system for auditing, debugging, and analytics.
type EventLog struct {
    gorm.Model
    Type      EventType `gorm:"type:varchar(100);not null;index"` // Type of event
    UserID    uint      `gorm:"index"`                             // Associated user, if applicable
    Details   string    `gorm:"type:text"`                         // Human-readable description or details of the event
    Metadata  string    `gorm:"type:text"`                         // JSON-encoded string for additional metadata about the event
    Timestamp time.Time `gorm:"index;not null"`                    // Timestamp of the event
}

// Key Enhancements:
// - **EventType Enumeration**: Clearly categorizes each log entry, facilitating easier filtering and analysis.
// - **Optional User Association**: Allows linking events to specific users where applicable, enhancing traceability.
// - **Detailed Metadata Storage**: Storing additional metadata as a JSON-encoded string provides flexibility to capture varied and complex information without altering the database schema.

// Implementing an `EventLog` model like this helps in creating a comprehensive log of significant actions and changes within your platform, serving a wide range of purposes from security auditing to user behavior analysis.
