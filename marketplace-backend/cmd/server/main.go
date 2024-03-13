package main

import (
    "log"
    "net/http"
    "os"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite" // Example: using SQLite; replace with your actual database driver

    "marketplace-backend/internal/api/v1"
    "marketplace-backend/internal/config"
    "marketplace-backend/internal/middleware"
    "marketplace-backend/internal/util"
)

func main() {
    // Load environment variables from .env file if present
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found or error loading .env file")
    }

    // Initialize the Gin router with default middleware
    r := gin.Default()

    // Apply middleware
    r.Use(gin.Logger())
    r.Use(middleware.RequestLogger()) // Log each request
    r.Use(gin.Recovery())
    r.Use(middleware.Recovery()) // Custom recovery to handle panics gracefully
    r.Use(middleware.SecureHeaders()) // Set secure headers
    r.Use(middleware.CORSMiddleware()) // Apply CORS settings

    // Rate limiter configuration - example: 1 request/second
    rl := middleware.NewRateLimiter(1, 5)
    r.Use(rl.Middleware())

    // Initialize database connection
    db, err := config.ConnectDatabase()
    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }


    // Optionally, apply JWT middleware globally or to specific routes
    // r.Use(middleware.JWTAuthentication())

    // Register API routes
    setupAPIRoutes(r, db)

    // Start the HTTP server with dynamic port configuration
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port if not specified
    }
    log.Printf("Starting server on port %s", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}

// initDB initializes the database connection using the configuration package
func initDB() (*gorm.DB, error) {
    dbURL := os.Getenv("DATABASE_URL")
    db, err := gorm.Open(sqlite.Open(dbURL), &gorm.Config{}) // Replace with actual DB config
    if err != nil {
        return nil, err
    }
    // AutoMigrate your models here (e.g., db.AutoMigrate(&model.User{}))
    return db, nil
}
    // Setup JWT Middleware if you have authentication
    // Example: r.Use(middleware.JWTAuthentication())

    // Start background worker for processing tasks
    go util.Worker()

    // Setup API routes
    setupAPIRoutes(r, db)

    // Retrieve the port from environment variables to allow flexibility
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port if not specified
    }

    // Start the HTTP server
    if err := r.Run(":" + port); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}

func setupAPIRoutes(r *gin.Engine, db *gorm.DB) {
    // Protect routes with JWT Authentication as needed
    authMiddleware := middleware.JWTAuthentication()

    api := r.Group("/api")
    {
        v1Group := api.Group("/v1", authMiddleware)
        {
            // Direct route registrations are removed in favor of modular registration functions
			v1Group.POST("/listings", v1.CreateListing(db))
			v1Group.GET("/listings", v1.GetAllListings(db))
			v1Group.GET("/user/listings", v1.GetUserListings(db))
			// Add routes for update and delete as needed
			v1Group.POST("/images/upload", v1.UploadImage(db))
			v1Group.GET("/listings/search", v1.SearchListings(db))
			v1Group.POST("/reviews", v1.PostReview(db))
			v1Group.GET("/users/:userID/reviews", v1.GetReviews(db))
			v1Group.POST("/messages/send", v1.SendMessage(db))
			v1Group.GET("/messages", v1.GetMessages(db))
			v1Group.POST("/transactions/initiate", v1.InitiateTransaction(db))
			v1Group.PATCH("/transactions/:transactionID/status", v1.UpdateTransactionStatus(db))
			v1Group.POST("/payments/initiate", v1.InitiatePayment)
			v1Group.GET("/analytics/seller", v1.GetSellerAnalytics(db))
			api.POST("/upload", v1.UploadFile)
			v1Group.POST("/tags/addToListing", v1.AddTagsToListing(db))
			v1Group.GET("/tags/:tagName/listings", v1.GetListingsByTag(db))
			v1Group.GET("/rewards", v1.GetUserRewards(db))
			v1Group.POST("/rewards/redeem", v1.RedeemReward(db))
        }

        // WebSocket route for real-time chat within the API routes setup for clarity
        v1Group.GET("/ws", func(c *gin.Context) {
            util.HandleChatConnections(c.Writer, c.Request)
        })
    }

    // Admin-specific routes protected by AdminRequired middleware
    adminRoutes := api.Group("/admin", middleware.AdminRequired())
    {
        adminRoutes.GET("/dashboard", v1.DashboardOverview(db))
        adminRoutes.GET("/reports", v1.GetReports(db)) // Example admin route to view reports
        // Additional admin routes as needed...
    }
}




























// package main

// import (
// 	"log"
// 	"marketplace-backend/internal/api/v1"
// 	"marketplace-backend/internal/config"
// 	"os"

// 	"github.com/gin-gonic/gin"
// 	"github.com/joho/godotenv"
// )

// func main() {
// 	// Load environment variables from .env file if present
// 	if err := godotenv.Load(); err != nil {
// 		log.Println("No .env file found")
// 	}
// 	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
// 		util.HandleConnections(w, r)
// 	})
// 	// Connect to the database
// 	db, err := config.ConnectDatabase()
// 	if err != nil {
// 		log.Fatalf("Could not connect to database: %v", err)
// 	}
// 	go util.Worker() // Start the worker in a non-blocking manner
// 	// Initialize the Gin router
// 	r := gin.Default()
// 	go util.StartMessageHandler() // Start handling chat messages

// 	r.GET("/ws", func(c *gin.Context) {
// 		util.HandleChatConnections(c.Writer, c.Request)
// 	})
// 	// Register the API endpoints
// 	v1.RegisterUserRoutes(r, db)
// 	api := r.Group("/api")
// 	{
// 		v1Group := api.Group("/v1")
// 		v1Group.Use(middleware.JWTAuthentication()) // Protect these routes
// 		{
// 			v1Group.POST("/listings", v1.CreateListing(db))
// 			v1Group.GET("/listings", v1.GetAllListings(db))
// 			v1Group.GET("/user/listings", v1.GetUserListings(db))
// 			// Add routes for update and delete as needed
// 			v1Group.POST("/images/upload", v1.UploadImage(db))
// 			v1Group.GET("/listings/search", v1.SearchListings(db))
// 			v1Group.POST("/reviews", v1.PostReview(db))
// 			v1Group.GET("/users/:userID/reviews", v1.GetReviews(db))
// 			v1Group.POST("/messages/send", v1.SendMessage(db))
// 			v1Group.GET("/messages", v1.GetMessages(db))
// 			v1Group.POST("/transactions/initiate", v1.InitiateTransaction(db))
// 			v1Group.PATCH("/transactions/:transactionID/status", v1.UpdateTransactionStatus(db))
// 			v1Group.POST("/payments/initiate", v1.InitiatePayment)
// 			v1Group.GET("/analytics/seller", v1.GetSellerAnalytics(db))
// 			api.POST("/upload", v1.UploadFile)
// 			v1Group.POST("/tags/addToListing", v1.AddTagsToListing(db))
// 			v1Group.GET("/tags/:tagName/listings", v1.GetListingsByTag(db))
// 			v1Group.GET("/rewards", v1.GetUserRewards(db))
// 			v1Group.POST("/rewards/redeem", v1.RedeemReward(db))

// 		}
// 	}
// 	adminRoutes := r.Group("/admin").Use(middleware.AdminRequired())
// 	{
// 		adminRoutes.GET("/dashboard", v1.DashboardOverview(db))
// 		// Additional admin routes...
// 	}

// 	return r
// }

// 	// Retrieve the port from environment variables to allow flexibility
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080" // Default port if not specified
// 	}

// 	// Start the HTTP server
// 	if err := r.Run(":" + port); err != nil {
// 		log.Fatalf("Failed to run server: %v", err)
// 	}
// }
// func setupRouter(db *gorm.DB) *gin.Engine {
// 	r := gin.Default()

// 	api := r.Group("/api")
// 	{
// 		v1Group := api.Group("/v1")
// 		v1Group.GET("/recommendations", v1.GetUserRecommendations(db))
// 	}

// 	return r
// }
// func setupRouter(db *gorm.DB) *gin.Engine {
// 	r := gin.Default()

// 	api := r.Group("/api")
// 	{
// 		v1Group := api.Group("/v1")
// 		{
// 			v1Group.POST("/reports", v1.SubmitReport(db))  // All users can submit reports
// 			adminGroup := v1Group.Group("/admin")
// 			adminGroup.Use(middleware.AdminRequired())     // Admin middleware to protect admin routes
// 			{
// 				adminGroup.GET("/reports", v1.GetReports(db)) // Only admins can view reports
// 			}
// 		}
// 	}

// 	return r
// }
// package main

// import (
//     "marketplace-backend/internal/model"
//     "marketplace-backend/internal/util"
//     // Import other necessary packages
// )

// func main() {
//     // Example trigger: user signs up for an event on the platform
//     user := model.User{
//         // User details...
//     }

//     // Assuming you have a function to check if the user opts in for marketing emails
//     if user.EmailPreferences.Marketing {
//         subject := "Thank you for signing up for the event!"
//         content := "<p>We're excited to have you join us. Here's what you need to know...</p>"
//         err := util.SendEmail(user.Email, subject, content)
//         if err != nil {
//             // Handle error
//         }
//     }

//     // Additional logic...
// }
// package main

// import (
// 	"marketplace-backend/internal/util"
// 	"fmt"
// )

// func main() {
// 	// Example scenario: User has earned a reward
// 	rewardTitle := "I just earned the Super Seller badge on Marketplace Platform!"
// 	rewardURL := "https://yourplatform.com/rewards/super-seller"

// 	shareURLs := util.GenerateSocialShareURLs(rewardTitle, rewardURL)

// 	fmt.Println("Share your achievement on social media:")
// 	for platform, url := range shareURLs {
// 		fmt.Printf("%s: %s\n", platform, url)
// 	}

// 	// Output these URLs in a user-facing context, such as part of a web response, email, or in-app notification.
// }
