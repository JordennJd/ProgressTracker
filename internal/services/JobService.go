package services

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"progress-tracker/internal/models"
	"progress-tracker/internal/queries"
	"time"
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
	err := s.db.Model(&models.Job{}).Where("job_id = ?", query.JobID).
		Updates(map[string]interface{}{
			"status":     models.StatusRunning,
			"started_at": time.Now(),
		}).Error

	return err
}

func (s *JobService) CompleteJob(query queries.CompleteJobQuery, userId uuid.UUID) error {
	if !s.IsJobExists(query.JobID) {
		return errors.New(fmt.Sprintf("Job with job id: %s doesn't exists", query.JobID))
	}

	if query.IsFailed {
		err := s.db.Model(&models.Job{}).Where("job_id = ?", query.JobID).
			Updates(map[string]interface{}{
				"status":      models.StatusFailed,
				"message":     query.Message,
				"finished_at": time.Now(),
			}).
			Error
		return err
	} else {
		err := s.db.Model(&models.Job{}).Where("job_id = ?", query.JobID).
			Updates(map[string]interface{}{
				"status":      models.StatusCompleted,
				"message":     query.Message,
				"finished_at": time.Now(),
			}).
			Error
		return err
	}
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

func (s *JobService) GetAll(dateFrom, dateTo time.Time) ([]models.Job, error) {
	var jobs []models.Job

	if dateTo.After(dateFrom) {
		return nil, fmt.Errorf("dateFrom cannot be after dateTo")
	}

	query := s.db.Raw(`
        SELECT * FROM job 
        WHERE created_at BETWEEN ? AND ?
        ORDER BY created_at ASC`,
		dateFrom.AddDate(0, 0, -1).Format("2006-01-02 15:04:05"),
		dateTo.AddDate(0, 0, 1).Format("2006-01-02 15:04:05"),
	)

	result := query.Scan(&jobs)
	if result.Error != nil {
		return nil, result.Error
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

func (s *JobService) GetNextJob(jobType string) (*models.Job, error) {
	var job models.Job
	err := s.db.Where("job_type = ? AND status = ?", jobType, models.StatusCreated).
		Order("created_at ASC").
		First(&job).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &job, nil
}

func (s *JobService) GetWorkingJobs(jobType string) ([]models.Job, error) {
	var jobs []models.Job
	err := s.db.Where("job_type = ? AND status = ?", jobType, models.StatusRunning).
		Find(&jobs).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return jobs, nil
}
