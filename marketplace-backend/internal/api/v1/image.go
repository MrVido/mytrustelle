package v1

import (
    "fmt"
    "io"
    "net/http"
    "strconv"

    "cloud.google.com/go/storage"
    "github.com/gin-gonic/gin"
    "google.golang.org/api/option"
    "marketplace-backend/internal/model"
    "gorm.io/gorm"
)

var (
    storageClient *storage.Client
    bucketName    = os.Getenv("GCS_BUCKET_NAME") // Ensure this environment variable is set.
)

func init() {
    // Initialize Google Cloud client
    ctx := context.Background()
    var err error
    storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("path/to/google-credentials.json")) // Specify your GCP credentials file.
    if err != nil {
        panic(fmt.Sprintf("Failed to create Google Cloud Storage client: %v", err))
    }
}

// UploadImage handles the uploading of images for a listing.
func UploadImage(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Retrieve the user ID from the JWT middleware (assumes middleware is set up correctly)
        userIDInterface, exists := c.Get("userID")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
            return
        }
        userID := userIDInterface.(uint) // Type assertion; make sure this is safe and handled properly.

        // Retrieve the listing ID from the form and validate it.
        listingIDStr := c.PostForm("listingID")
        listingID, err := strconv.Atoi(listingIDStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid listing ID"})
            return
        }

        // Retrieve the image file from the form.
        file, header, err := c.Request.FormFile("image")
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Must provide an image file"})
            return
        }
        defer file.Close()

        // Validate the file size and type if needed.

        // Upload the file to Google Cloud Storage.
        objectName := fmt.Sprintf("images/%d/%s", userID, header.Filename) // Include the user ID in the object path for organization and to prevent filename collisions.
        ctx := context.Background()
        wc := storageClient.Bucket(bucketName).Object(objectName).NewWriter(ctx)
        if _, err = io.Copy(wc, file); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image to cloud storage"})
            return
        }
        if err = wc.Close(); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to finalize image upload to cloud storage"})
            return
        }

        // Set the image to be publicly readable.
        if err := storageClient.Bucket(bucketName).Object(objectName).ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set image ACL to public"})
            return
        }

        // Retrieve the URL of the uploaded object.
        attrs, err := storageClient.Bucket(bucketName).Object(objectName).Attrs(ctx)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get image attributes"})
            return
        }

        // Create and save the image record in the database.
        image := model.Image{ListingID: uint(listingID), URL: attrs.MediaLink}
        if result := db.Create(&image); result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image record in database"})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"message": "Image uploaded successfully", "imageURL": attrs.MediaLink})
    }
}
