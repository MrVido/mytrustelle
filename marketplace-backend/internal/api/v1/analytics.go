package v1

import (
    "marketplace-backend/internal/model"
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// GetSellerAnalytics provides an overview of a seller's performance.
func GetSellerAnalytics(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Assuming userID is extracted from JWT token and validated
        userID, exists := c.Get("userID")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
            return
        }

        var totalSales float64
        var transactionsCount int64
        var averageRating float64
        var totalListingViews int64

        // Example query to calculate total sales from transactions table for the seller
        db.Table("transactions").Select("sum(amount)").Where("seller_id = ?", userID).Row().Scan(&totalSales)

        // Count the number of transactions
        db.Model(&model.Transaction{}).Where("seller_id = ?", userID).Count(&transactionsCount)

        // Calculate average rating from reviews table for the seller
        db.Table("reviews").Select("avg(rating)").Where("seller_id = ?", userID).Row().Scan(&averageRating)

        // For totalListingViews, you'll need to track views on listings. This is just a placeholder.
        // Assuming there's a "views" field in your listings table or a separate table to track views.
        db.Model(&model.Listing{}).Select("sum(views)").Where("seller_id = ?", userID).Row().Scan(&totalListingViews)

        analyticsData := map[string]interface{}{
            "totalSales":         totalSales,
            "transactionsCount":  transactionsCount,
            "averageRating":      averageRating,
            "totalListingViews":  totalListingViews,
        }

        c.JSON(http.StatusOK, analyticsData)
    }
}
