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

        // Pagination
        page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
        pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
        offset := (page - 1) * pageSize

        // Finalize the query with pagination
        query = query.Offset(offset).Limit(pageSize)

        // Execute the query
        if result := query.Find(&listings); result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
            return
        }

        c.JSON(http.StatusOK, listings)
    }
}

// CreateListing - Added casting safety check and additional validation
func CreateListing(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var listing model.Listing
        if err := c.ShouldBindJSON(&listing); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid listing details"})
            return
        }

        userID, exists := c.Get("userID")
        if !exists || userID == nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }

        // Additional validation can be added here (e.g., required fields, field lengths)

        listingUserID, ok := userID.(uint)
        if !ok {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID conversion error"})
            return
        }

        listing.UserID = listingUserID

        if result := db.Create(&listing); result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
            return
        }

        c.JSON(http.StatusCreated, listing)
    }
}

// GetAllListings returns all listings available in the marketplace.
// GetListing retrieves a specific listing, potentially from cache for performance.
func GetListing(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        listingID := c.Param("id")
        var listing model.Listing

        // Attempt to retrieve the listing from the cache first
        cacheKey := "listing:" + listingID
        if err := util.GetCache(cacheKey, &listing); err == nil {
            c.JSON(http.StatusOK, listing)
            return
        }

        // Cache miss; fetch the listing from the database and cache it
        if err := db.First(&listing, listingID).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Listing not found"})
            return
        }
        util.SetCache(cacheKey, listing, 30*time.Minute)

        c.JSON(http.StatusOK, listing)
    }
}


// GetUserListings returns listings created by the logged-in user.
func GetUserListings(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID, exists := c.Get("userID")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }

        var listings []model.Listing
        page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
        pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
        offset := (page - 1) * pageSize

        db.Where("user_id = ?", userID).Offset(offset).Limit(pageSize).Find(&listings)
        c.JSON(http.StatusOK, listings)
    }
}