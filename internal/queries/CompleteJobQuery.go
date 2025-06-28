package queries

import "github.com/google/uuid"

type CompleteJobQuery struct {
	JobID    uuid.UUID `json:"job_id"`
	Message  *string   `json:"message"`
	IsFailed bool      `json:"is_failed"`
}
