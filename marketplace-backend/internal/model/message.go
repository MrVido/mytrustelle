package model

import (
	"gorm.io/gorm"
)

// Message represents a message between users, incorporating advanced features.
type Message struct {
	gorm.Model
	SenderID   uint           `gorm:"index;not null"` // Indexing for faster queries based on SenderID
	ReceiverID uint           `gorm:"index;not null"` // Indexing for faster queries based on ReceiverID
	Content    string         `gorm:"type:text;not null"` // The content of the message
	CreatedAt  time.Time     `gorm:"index"` // Indexing to allow sorting by creation time
	ReadAt     *time.Time     // Tracks when the message was read. Nil if unread.
	ConversationID uint       `gorm:"index;not null"` // Groups messages into conversations
	DeletedAt  gorm.DeletedAt `gorm:"index"` // Soft delete to hide messages without permanent deletion
}

// Conversation represents a collection of messages between two users.
type Conversation struct {
	gorm.Model
	Participant1 uint `gorm:"index"` // First participant's user ID
	Participant2 uint `gorm:"index"` // Second participant's user ID
	LastMessageID uint // ID of the last message in the conversation, for quick access
	LastMessagePreview string // Optional: Store a snippet or the full content of the last message for display in conversation list
	DeletedAt    gorm.DeletedAt `gorm:"index"` // Soft delete for the entire conversation
}

// Key Enhancements:

// 1. **Read Receipts**: The `ReadAt` field allows tracking of when a message was read by the recipient, enabling read receipts functionality.

// 2. **Message Threading**: `ConversationID` associates each message with a conversation thread, allowing efficient retrieval and display of related messages.

// 3. **Soft Deletes**: Both messages and conversations support soft deletion, enabling users to "delete" them without removing the data from the database, which is crucial for dispute resolution or moderation.

// 4. **Performance**: Indexes on `SenderID`, `ReceiverID`, `CreatedAt`, and `ConversationID` ensure efficient query operations, crucial for a system with a potentially high volume of messages.

// 5. **Conversation Management**: The `Conversation` model enables tracking and management of conversations between users, including quick access to the last message and soft deletion of conversations.

// These enhancements ensure that your messaging system is robust, user-friendly, and ready to handle the dynamic communication needs of a marketplace platform.
