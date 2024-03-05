package model

import (
	"gorm.io/gorm"
	"time"
)

// ChatMessage represents a message in a chat between two users.
type ChatMessage struct {
	gorm.Model
	SenderID    uint      `gorm:"not null"` // ID of the user sending the message
	ReceiverID  uint      `gorm:"not null"` // ID of the user receiving the message
	Content     string    `gorm:"type:text;not null"`
	SentAt      time.Time
}
