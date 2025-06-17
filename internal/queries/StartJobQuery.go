package queries

import "github.com/google/uuid"

type StartJobQuery struct {
	JobID uuid.UUID `json:"job_id"`
}
