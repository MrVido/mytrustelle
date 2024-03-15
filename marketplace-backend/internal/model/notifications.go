package model

import (
	"gorm.io/gorm"
	"time"
)

// NotificationType categorizes the type of notification
type NotificationType string

const (
	NotificationTypeMessage NotificationType = "message"
	NotificationTypeOffer   NotificationType = "offer"
	NotificationTypeSystem  NotificationType = "system"
	// Add more types as needed
)

// Notification represents a user notification for various events on the platform.
type Notification struct {
	gorm.Model
	UserID     uint             `gorm:"index;not null"` // Target user for the notification
	Type       NotificationType `gorm:"type:varchar(20);not null"`
	Title      string           // Brief title for the notification
	Message    string           `gorm:"type:text;not null"` // Main notification content
	Read       bool             `gorm:"default:false"` // Whether the notification has been read
	ActionURL  string           // Optional URL for a call-to-action
	CreatedAt  time.Time
}

// Key Enhancements:
// - **NotificationType**: Categorizes notifications to allow users to filter or prioritize them.
// - **ActionURL**: Provides a direct link to the relevant section/item on the platform, improving user engagement.
// - **Read Status Tracking**: Helps in managing unread notifications, enhancing user interaction with the platform.
