package v1

import (
	"marketplace-backend/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DashboardOverview provides a summary of platform activity for the admin.
func DashboardOverview(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Example: Count users, transactions, and active listings
		var userCount, transactionCount, listingCount int64
		db.Model(&model.User{}).Count(&userCount)
		db.Model(&model.Transaction{}).Count(&transactionCount)
		db.Model(&model.Listing{}).Count(&listingCount)

		// Example response. Expand with more detailed analytics as needed.
		c.JSON(http.StatusOK, gin.H{
			"userCount":        userCount,
			"transactionCount": transactionCount,
			"listingCount":     listingCount,
		})
	}
}
