package v1

import (
	"marketplace-backend/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SendMessage allows a user to send a message to another user.
func SendMessage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var message model.Message
		if err := c.ShouldBindJSON(&message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Ensure the sender is not the same as the receiver
		if message.SenderID == message.ReceiverID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Sender and receiver must be different"})
			return
		}

		if result := db.Create(&message); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Message sent successfully"})
	}
}

// GetMessages retrieves messages between two users.
func GetMessages(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		senderID := c.Query("senderID")
		receiverID := c.Query("receiverID")

		var messages []model.Message
		db.Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
			senderID, receiverID, receiverID, senderID).Find(&messages)

		c.JSON(http.StatusOK, messages)
	}
}
