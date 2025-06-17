package queries

import "github.com/google/uuid"

type SetProgressQuery struct {
	JobID    uuid.UUID `json:"job_id"`
	Progress float32   `json:"progress"`
}
