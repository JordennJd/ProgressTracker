package queries

import "github.com/google/uuid"

type CreateJobQuery struct {
	JobID   uuid.UUID `json:"job_id"`
	JobType string    `json:"job_type"`
	Message *string   `json:"message"`
	Context *string   `json:"context"`
}
