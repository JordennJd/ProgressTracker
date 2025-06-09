package models

import (
	"time"

	"github.com/google/uuid"
)

type JobStatus int

const (
	StatusPending   JobStatus = iota // 0
	StatusRunning                    // 1
	StatusCompleted                  // 2
	StatusFailed                     // 3
)

type Job struct {
	ID         uuid.UUID  `json:"id"`
	JobID      uuid.UUID  `json:"job_id"`
	Status     JobStatus  `json:"status"`
	Progress   int        `json:"progress"`
	JobType    string     `json:"job_type"`
	Message    *string    `json:"message"`
	Context    *string    `json:"context"`
	CreatedAt  time.Time  `json:"created_at"`
	FinishedAt *time.Time `json:"finished_at"`
}

func (Job) TableName() string {
	return "job"
}
