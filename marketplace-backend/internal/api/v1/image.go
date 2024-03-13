package v1

import (
    "fmt"
    "io/ioutil"
    "mime/multipart"
    "net/http"
    "os"
    "strconv"

    "github.com/gin-gonic/gin"
    "golang.org/x/net/context"
    "golang.org/x/oauth2/google"
    "cloud.google.com/go/storage"
    "marketplace-backend/internal/model"
    "gorm.io/gorm"
)

// Assuming you have set up Google Cloud credentials and bucket name as environment variables
var (
    storageClient *storage.Client
    bucketName    = os.Getenv("GCS_BUCKET_NAME")
)

func init() {
    ctx := context.Background()
    client, err := google.DefaultClient(ctx, storage.ScopeFullControl)
    if err != nil {
        panic(fmt.Sprintf("Failed to create Google Cloud client: %v", err))
    }
    storageClient, err = storage.NewClient(ctx, option.WithHTTPClient(client))
    if err != nil {
        panic(fmt.Sprintf("Failed to create Google Cloud Storage client: %v", err))
    }
}

// uploadToCloudStorage uploads a file to Google Cloud Storage and returns the file URL.
func uploadToCloudStorage(file multipart.File, header *multipart.FileHeader) (string, error) {
    ctx := context.Background()
    bucket := storageClient.Bucket(bucketName)
    
    // Ensure unique file names in the bucket, could use UUIDs or similar approach
    objectName := fmt.Sprintf("images/%s", header.Filename)
    obj := bucket.Object(objectName)
    
    w := obj.NewWriter(ctx)
    if _, err := ioutil.ReadAll(file); err != nil {
        return "", err
    }
    if err := w.Close(); err != nil {
        return "", err
    }
    
    if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
        return "", err
    }

    attrs, err := obj.Attrs(ctx)
    if err != nil {
        return "", err
    }

    return attrs.MediaLink, nil
}

// UploadImage handles the uploading of images for a listing.
func UploadImage(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID, _ := c.Get("userID") // Assumes userID is extracted from JWT token
        listingID, err := strconv.Atoi(c.PostForm("listingID"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid listing ID"})
            return
        }

        file, header, err := c.Request.FormFile("image")
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Must provide an image file"})
            return
        }
        
        // Validate file size and type if needed here
        
        // For cloud storage (e.g., Google Cloud Storage)
        uploadedURL, err := uploadToCloudStorage(file, header)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image to cloud storage"})
            return
        }

        // Create and save the image record in the database
        image := model.Image{ListingID: uint(listingID), URL: uploadedURL}
        if result := db.Create(&image); result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image record in database"})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"message": "Image uploaded successfully", "imageURL": uploadedURL})
    }
}
