package model

import (
	"gorm.io/gorm"
	"time"
)

// ChatMessage represents a message in a chat between two users, incorporating soft deletes.
type ChatMessage struct {
	gorm.Model
	SenderID    uint      `gorm:"index;not null"` // ID of the user sending the message, indexed for faster queries
	ReceiverID  uint      `gorm:"index;not null"` // ID of the user receiving the message, indexed for faster queries
	Content     string    `gorm:"type:text;not null"` // The content of the message
	SentAt      time.Time `gorm:"index"` // Indexed to allow efficient sorting and querying by time
	Read        bool      `gorm:"default:false"` // Indicates whether the message has been read by the receiver
	DeletedAt   gorm.DeletedAt `gorm:"index"` // Soft delete field, allows hiding messages without permanent deletion
}

// Chat represents a chat session between two users, also incorporating soft deletes.
type Chat struct {
	gorm.Model
	Participant1 uint      `gorm:"index"` // ID of one participant in the chat
	Participant2 uint      `gorm:"index"` // ID of the other participant in the chat
	LastMessage  string    `gorm:"type:text"` // Snapshot of the last message sent in this chat
	LastSentAt   time.Time `gorm:"index"` // Time of the last message, indexed for sorting chats by recent activity
	DeletedAt    gorm.DeletedAt `gorm:"index"` // Soft delete field, allows hiding entire chat sessions without permanent deletion
}

// Timestamp management: Ensure that SentAt and LastSentAt are automatically updated to the current time when new messages are sent.
// This can be achieved in GORM by using hooks or callback methods that GORM provides, such as BeforeCreate or BeforeUpdate.
