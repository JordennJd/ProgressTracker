package services

import (
	"errors"
	"fmt"
	"progress-tracker/internal/models"
	"progress-tracker/internal/queries"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobService struct {
	db *gorm.DB
}

func NewJobService(db *gorm.DB) *JobService {
	return &JobService{db: db}
}

func (s *JobService) CreateJob(query queries.CreateJobQuery, userId uuid.UUID) error {
	if s.IsJobExists(query.JobID) {
		return errors.New(fmt.Sprintf("Job with job id: %s already created", query.JobID))
	}

	model := models.Job{
		ID:        uuid.New(),
		JobID:     query.JobID,
		UserID:    userId,
		Status:    models.StatusCreated,
		JobType:   query.JobType,
		Context:   query.Context,
		Message:   query.Message,
		CreatedAt: time.Now(),
		Progress:  0,
	}

	err := s.db.Create(model).Error

	return err
}

func (s *JobService) StartJob(query queries.StartJobQuery, userId uuid.UUID) error {
	if !s.IsJobExists(query.JobID) {
		return errors.New(fmt.Sprintf("Job with job id: %s doesn't exists", query.JobID))
	}
	err := s.db.Model(&models.Job{}).Where("job_id = ?", query.JobID).Update("status", models.StatusRunning).Error

	return err
}

func (s *JobService) GetJobByID(id uuid.UUID) (*models.Job, error) {
	var job models.Job
	result := s.db.Raw("SELECT * FROM job WHERE id = ?", id).Scan(&job)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &job, nil
}

func (s *JobService) GetJobByJobID(jobId uuid.UUID) (*models.Job, error) {
	var job models.Job
	result := s.db.Raw("SELECT * FROM job WHERE job_id = ?", jobId).Scan(&job)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &job, nil
}

func (s *JobService) GetAll() ([]models.Job, error) {
	var jobs []models.Job
	result := s.db.Raw("SELECT * FROM job").Scan(&jobs)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return jobs, nil
}

func (s *JobService) IsJobExists(jobId uuid.UUID) bool {
	job, err := s.GetJobByJobID(jobId)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	return job != nil
}
