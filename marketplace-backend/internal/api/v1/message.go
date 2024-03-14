package v1

import (
	"marketplace-backend/internal/model"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SendMessage allows a user to send a message to another user, ensuring sender authentication.
func SendMessage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var message model.Message
		if err := c.ShouldBindJSON(&message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message format"})
			return
		}

		// Authentication: Retrieve the sender's ID from the JWT token to ensure sender authenticity.
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Ensure the authenticated user is the sender.
		if message.SenderID != userID.(uint) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Sender ID does not match authenticated user"})
			return
		}

		// Prevent sending messages to oneself.
		if message.SenderID == message.ReceiverID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Sender and receiver must be different"})
			return
		}

		// Create the message in the database.
		if result := db.Create(&message); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Message sent successfully", "data": message})
	}
}

// GetMessages retrieves messages between two users, ensuring one of them is the authenticated user.
func GetMessages(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		senderID := c.Query("senderID")
		receiverID := c.Query("receiverID")

		// Check if the authenticated user is part of the conversation.
		if userID != senderID && userID != receiverID {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access to the requested messages is forbidden"})
			return
		}

		var messages []model.Message
		db.Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
			senderID, receiverID, receiverID, senderID).Order("created_at desc").Find(&messages)

		c.JSON(http.StatusOK, messages)
	}
}
