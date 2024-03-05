package v1

import (
	"marketplace-backend/internal/model"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// InitiateTransaction starts a new transaction for a listing.
func InitiateTransaction(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var transaction model.Transaction
		if err := c.ShouldBindJSON(&transaction); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		transaction.Status = model.TransactionPending

		if result := db.Create(&transaction); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusCreated, transaction)
	}
}

// UpdateTransactionStatus changes the status of an existing transaction.
func UpdateTransactionStatus(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		transactionID := c.Param("transactionID")
		var update struct {
			Status model.TransactionStatus `json:"status"`
		}
		if err := c.ShouldBindJSON(&update); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if result := db.Model(&model.Transaction{}).Where("id = ?", transactionID).Update("status", update.Status); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Transaction status updated successfully"})
	}
}
