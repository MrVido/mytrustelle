package v1

import (
    "marketplace-backend/internal/model"
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// InitiateTransaction starts a new transaction for a listing with added security and validation.
func InitiateTransaction(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var transaction model.Transaction
        if err := c.ShouldBindJSON(&transaction); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction format"})
            return
        }

        // Ensure the listing exists
        var listing model.Listing
        if err := db.First(&listing, transaction.ListingID).Error; err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Listing not found"})
            return
        }

        // Check if the user is authorized to initiate the transaction
        userID, exists := c.Get("userID")
        if !exists || userID != transaction.BuyerID {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized or invalid buyer ID"})
            return
        }

        transaction.Status = model.TransactionPending
        if result := db.Create(&transaction); result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initiate transaction"})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"message": "Transaction initiated successfully", "transaction": transaction})
    }
}

// UpdateTransactionStatus changes the status of an existing transaction with authorization checks.
func UpdateTransactionStatus(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        transactionID := c.Param("transactionID")
        var update struct {
            Status model.TransactionStatus `json:"status"`
        }
        if err := c.ShouldBindJSON(&update); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status update request"})
            return
        }

        // Retrieve the transaction to check if the user is authorized to update it
        var transaction model.Transaction
        if err := db.First(&transaction, transactionID).Error; err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Transaction not found"})
            return
        }

        // Authorization: Ensure the user updating the transaction is involved (as a buyer or seller)
        userID, exists := c.Get("userID")
        if !exists || (userID != transaction.BuyerID && userID != transaction.SellerID) {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized to update this transaction"})
            return
        }

        // Update the transaction status
        if result := db.Model(&transaction).Update("status", update.Status); result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update transaction status"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Transaction status updated successfully"})
    }
}
