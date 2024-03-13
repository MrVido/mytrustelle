package config

import (
    "context"
    "os"
    "strconv"
    "log"

    "github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

// NewRedisClient creates and configures an instance of the Redis client.
func NewRedisClient() *redis.Client {
    // Load Redis configuration from environment variables
    addr := os.Getenv("REDIS_ADDR")
    if addr == "" {
        addr = "localhost:6379" // Default address
    }

    password := os.Getenv("REDIS_PASSWORD") // Default is no password
    dbStr := os.Getenv("REDIS_DB")
    db, err := strconv.Atoi(dbStr) // Convert DB index from string to int
    if err != nil {
        db = 0 // Default to DB 0 if conversion fails or REDIS_DB is not set
        log.Printf("Invalid or no REDIS_DB provided, defaulting to 0: %v\n", err)
    }

    client := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password, // no password set by default
        DB:       db,       // default DB
        PoolSize: 100,      // Maximum number of socket connections.
        MinIdleConns: 10,   // Minimum number of idle connections which is useful for applications that open and close a lot of connections but always need a few idle.
    })

    // Ping the Redis server to check the connection
    _, err = client.Ping(Ctx).Result()
    if err != nil {
        log.Fatalf("Unable to connect to Redis: %v", err)
    }

    log.Println("Successfully connected to Redis")
    return client
}
