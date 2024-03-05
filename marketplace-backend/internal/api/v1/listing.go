package v1

import (
    "marketplace-backend/internal/model"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// SearchListings handles searching and filtering of listings.
func SearchListings(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var listings []model.Listing
        query := db

        // Filtering by category if provided
        if category := c.Query("category"); category != "" {
            query = query.Where("category = ?", category)
        }

        // Search in titles if search query provided
        if searchQuery := c.Query("q"); searchQuery != "" {
            likeQuery := "%" + searchQuery + "%"
            query = query.Where("title LIKE ?", likeQuery)
        }

        // Price range filtering
        if minPrice := c.Query("minPrice"); minPrice != "" {
            if min, err := strconv.ParseFloat(minPrice, 64); err == nil {
                query = query.Where("price >= ?", min)
            }
        }

        if maxPrice := c.Query("maxPrice"); maxPrice != "" {
            if max, err := strconv.ParseFloat(maxPrice, 64); err == nil {
                query = query.Where("price <= ?", max)
            }
        }

        // Execute the query
        if result := query.Find(&listings); result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
            return
        }

        c.JSON(http.StatusOK, listings)
    }
}


// CreateListing handles the creation of a new listing.
func CreateListing(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var listing model.Listing
		if err := c.ShouldBindJSON(&listing); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Assuming userID is set from JWT middleware
		userID, _ := c.Get("userID")
		listing.UserID = userID.(uint)

		if result := db.Create(&listing); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusCreated, listing)
	}
}
// GetAllListings returns all listings available in the marketplace.
func GetListing(db *gorm.DB, listingID string) (*model.Listing, error) {
	var listing model.Listing

	// Attempt to retrieve the listing from cache
	cacheKey := "listing:" + listingID
	if err := util.GetCache(cacheKey, &listing); err == nil && listing.ID != 0 {
		return &listing, nil // Return the cached listing
	}

	// Cache miss; fetch from database
	if err := db.First(&listing, listingID).Error; err != nil {
		return nil, err
	}

	// Cache the listing for future requests
	util.SetCache(cacheKey, listing, 30*time.Minute)

	return &listing, nil
}


// GetUserListings returns listings created by the logged-in user.
func GetUserListings(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")
		var listings []model.Listing
		db.Where("user_id = ?", userID).Find(&listings)
		c.JSON(http.StatusOK, listings)
	}
}
