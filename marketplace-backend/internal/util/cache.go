package util

import (
	"context"
	"encoding/json"
	"errors"
	"marketplace-backend/internal/config"
	"time"

	"github.com/go-redis/redis/v8"
)

// ctx should ideally be passed from higher-level functions to control lifecycle,
// but for simplicity, we'll define it globally here.
var ctx = context.Background()

// redisClient is a singleton instance of the Redis client to avoid reconnecting to Redis on each call.
var redisClient *redis.Client

// init initializes the Redis client. It's called when the package is imported.
func init() {
	redisClient = config.NewRedisClient()
}

// SetCache sets a value in Redis cache with a specified expiration time.
func SetCache(key string, value interface{}, expiration time.Duration) error {
	// Serialize the value to JSON
	serializedValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	// Set the key-value pair in Redis
	if err := redisClient.Set(ctx, key, serializedValue, expiration).Err(); err != nil {
		return err
	}

	return nil
}

// GetCache retrieves a value from Redis cache and decodes it into the specified object.
func GetCache(key string, dest interface{}) error {
	// Get the value from Redis
	val, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return errors.New("cache miss")
		}
		return err
	}

	// Deserialize the value from JSON
	if err := json.Unmarshal([]byte(val), dest); err != nil {
		return err
	}

	return nil
}
