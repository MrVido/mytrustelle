// GetSellerAnalytics provides an overview of a seller's performance.
func GetSellerAnalytics(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID, exists := c.Get("userID")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
            return
        }

        var totalSales float64
        var transactionsCount int64
        var averageRating float64
        var totalListingViews int64

        // Calculate total sales
        if err := db.Table("transactions").Select("sum(amount)").Where("seller_id = ?", userID).Row().Scan(&totalSales); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to calculate total sales"})
            return
        }

        // Count transactions
        if err := db.Model(&model.Transaction{}).Where("seller_id = ?", userID).Count(&transactionsCount).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to count transactions"})
            return
        }

        // Calculate average rating
        if err := db.Table("reviews").Select("avg(rating)").Where("seller_id = ?", userID).Row().Scan(&averageRating); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to calculate average rating"})
            return
        }

        // Calculate total listing views
        if err := db.Model(&model.Listing{}).Select("sum(views)").Where("seller_id = ?", userID).Row().Scan(&totalListingViews); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to calculate total listing views"})
            return
        }

        // You may add more detailed analytics here, such as sales trends or customer demographics

        c.JSON(http.StatusOK, map[string]interface{}{
            "totalSales":         totalSales,
            "transactionsCount":  transactionsCount,
            "averageRating":      averageRating,
            "totalListingViews":  totalListingViews,
        })
    }
}
