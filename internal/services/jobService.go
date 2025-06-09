package services

import (
	"progress-tracker/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobService struct {
	db *gorm.DB
}

func NewJobService(db *gorm.DB) *JobService {
	return &JobService{db: db}
}

func (s *JobService) SaveJob(job *models.Job) error {
	job.ID = uuid.New()
	return s.db.Create(job).Error
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
