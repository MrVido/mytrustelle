package model

import (
	"gorm.io/gorm"
	"time"
)

// Message represents a message between users.
type Message struct {
	gorm.Model
	SenderID   uint      `gorm:"not null"` // ID of the user sending the message
	ReceiverID uint      `gorm:"not null"` // ID of the user receiving the message
	Content    string    `gorm:"type:text;not null"`
	CreatedAt  time.Time
}
