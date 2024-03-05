package util

import (
	"context"
	"encoding/json"
	"marketplace-backend/internal/config"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// SetCache sets a value in Redis cache with a specified expiration time.
func SetCache(key string, value interface{}, expiration time.Duration) error {
	client := config.NewRedisClient()

	// Serialize the value to JSON
	serializedValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	// Set the key-value pair in Redis
	return client.Set(ctx, key, serializedValue, expiration).Err()
}

// GetCache retrieves a value from Redis cache and decodes it into the specified object.
func GetCache(key string, dest interface{}) error {
	client := config.NewRedisClient()

	// Get the value from Redis
	val, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil // Key does not exist
	} else if err != nil {
		return err
	}

	// Deserialize the value from JSON
	return json.Unmarshal([]byte(val), dest)
}
