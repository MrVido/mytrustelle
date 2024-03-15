package util

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"marketplace-backend/internal/config"
	"time"

	"github.com/go-redis/redis/v8"
)

type Job struct {
	Type    string                 `json:"type"`
	Payload map[string]interface{} `json:"payload"`
}

var ctx = context.Background()
var redisClient *redis.Client

func init() {
	redisClient = config.NewRedisClient()
}

// Enqueue adds a new job to the queue with improved error handling.
func Enqueue(job *Job) error {
	serializedJob, err := json.Marshal(job)
	if err != nil {
		return err
	}

	if err := redisClient.LPush(ctx, "jobs", serializedJob).Err(); err != nil {
		return err
	}

	return nil
}

// Worker listens for and processes jobs from the queue with structured error handling.
func Worker() {
	for {
		jobData, err := redisClient.BRPop(ctx, 0, "jobs").Result()
		if err != nil {
			if err == redis.Nil {
				time.Sleep(1 * time.Second) // Sleep for a short time if the queue is empty
				continue
			}
			log.Printf("Error retrieving job from queue: %v\n", err)
			continue
		}

		var job Job
		if err := json.Unmarshal([]byte(jobData[1]), &job); err != nil {
			log.Printf("Error unmarshalling job data: %v\n", err)
			continue
		}

		if err := processJob(&job); err != nil {
			log.Printf("Error processing job: %v\n", err)
			// Optionally, re-enqueue the job with a delay or move it to a dead-letter queue
		}
	}
}

// processJob handles job processing based on its type.
func processJob(job *Job) error {
	switch job.Type {
	case "send_email":
		// Replace with actual email sending logic
		return SendEmail(job)
	// Add more case statements for other job types
	default:
		return errors.New("unknown job type")
	}
}

// Placeholder for the SendEmail function used in job processing
func SendEmail(job *Job) error {
	// Email sending logic here
	log.Printf("Sending email: %+v\n", job)
	return nil
}
