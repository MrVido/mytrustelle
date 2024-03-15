package model

import (
	"gorm.io/gorm"
	"time"
)

// EventType categorizes the type of events being logged
type EventType string

const (
	UserLoginEvent     EventType = "USER_LOGIN"
	UserLogoutEvent    EventType = "USER_LOGOUT"
	TransactionEvent   EventType = "TRANSACTION"
	ListingViewEvent   EventType = "LISTING_VIEW"
	ErrorEvent         EventType = "ERROR"
	// Extend with more event types as necessary
)

// EventLog represents a record of significant events within the system.
type EventLog struct {
	gorm.Model
	Type        EventType `gorm:"type:varchar(100);not null;index"` // The category of the event
	UserID      uint      `gorm:"index;default:null"`               // User associated with the event, if applicable
	Description string    `gorm:"type:text;not null"`               // Human-readable description of the event
	Timestamp   time.Time `gorm:"default:CURRENT_TIMESTAMP"`        // Time the event occurred
}

// Enhancements and considerations directly included:
// 1. **Indexing**: Adding an index on the `Type` and optionally on `UserID` fields for faster querying of specific event types or user-related events.
// 2. **Event Types**: Defined as constants for type safety and ease of use, making it clear which events are being logged.
// 3. **User Association**: Optional association with a `UserID` allows for tracking events related to specific users, enhancing understanding of user behavior and troubleshooting user-reported issues.
// 4. **Timestamps**: Automatic timestamping of each event to record when it occurred, crucial for chronological event tracking and analysis.

