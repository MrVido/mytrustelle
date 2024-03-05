package util

import (
	"encoding/json"
	"marketplace-backend/internal/config"
	"github.com/go-redis/redis/v8"
)

type Job struct {
	Type    string                 `json:"type"`    // Job type, e.g., "send_email"
	Payload map[string]interface{} `json:"payload"` // Job payload
}

// Enqueue adds a new job to the queue.
func Enqueue(job *Job) error {
	serializedJob, err := json.Marshal(job)
	if err != nil {
		return err
	}

	client := config.NewRedisClient()
	err = client.LPush(config.Ctx, "jobs", serializedJob).Err()
	return err
}

// Worker listens for and processes jobs from the queue.
func Worker() {
	client := config.NewRedisClient()

	for {
		// Block until a job is available in the queue
		jobData, err := client.BRPop(config.Ctx, 0, "jobs").Result()
		if err != nil {
			continue // log error or handle as needed
		}

		var job Job
		if err := json.Unmarshal([]byte(jobData[1]), &job); err != nil {
			continue // log error or handle as needed
		}

		// Process the job based on its type
		switch job.Type {
		case "send_email":
			// process sending email
		// Add more case statements for other job types
		}
	}
}
