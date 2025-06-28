package models

import (
	"time"

	"github.com/google/uuid"
)

type JobStatus int

const (
	StatusCreated JobStatus = iota
	StatusPending
	StatusRunning
	StatusCompleted
	StatusFailed
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
	StartedAt  *time.Time `json:"started_at"`
	FinishedAt *time.Time `json:"finished_at"`
	UserID     uuid.UUID  `json:"user_id"`
}

func (Job) TableName() string {
	return "job"
}

func (j Job) GetMessage() string {
	if j.Message != nil {
		return *j.Message
	}
	return ""
}
