package handlers

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"progress-tracker/api/proto/models"
	pb "progress-tracker/api/proto/service"
	interceptors "progress-tracker/cmd/GRPC/interceprtors"
	"progress-tracker/internal/queries"
	"progress-tracker/internal/services"
)

type JobRpcServer struct {
	*services.JobService
	pb.UnimplementedJobServiceServer
}

func (j *JobRpcServer) GetWorkingJobs(ctx context.Context, in *models.GetWorkingJobRequest) (*models.GetJobsResult, error) {
	jobs, err := j.JobService.GetWorkingJobs(in.GetJobType())
	if err != nil {
		return &models.GetJobsResult{
			ErrorMessage: err.Error(),
			IsSuccessful: false,
			Jobs:         []*models.Job{},
		}, err
	}

	if jobs == nil {
		return &models.GetJobsResult{
			Jobs:         []*models.Job{},
			IsSuccessful: true,
		}, nil
	}

	protoJobs := make([]*models.Job, len(jobs))

	for index, job := range jobs {
		protoJobs[index] = &models.Job{
			Id:        job.ID.String(),
			JobId:     job.JobID.String(),
			UserId:    job.UserID.String(),
			Status:    models.JobStatus(job.Status),
			Message:   &wrappers.StringValue{Value: job.GetMessage()},
			JobType:   job.JobType,
			CreatedAt: timestamppb.New(job.CreatedAt),
			FinishedAt: func() *timestamppb.Timestamp {
				if job.FinishedAt != nil {
					return timestamppb.New(*job.FinishedAt)
				}
				return nil
			}(),
			Progress: int32(job.Progress),
		}
	}

	return &models.GetJobsResult{
		Jobs:         protoJobs,
		IsSuccessful: true,
	}, nil
}

func NewJobRpcServer(service *services.JobService) *JobRpcServer {
	return &JobRpcServer{
		JobService: service,
	}
}

func (j *JobRpcServer) GetNextJob(ctx context.Context, in *models.GetNextJobRequest) (*models.GetJobResult, error) {
	job, err := j.JobService.GetNextJob(in.GetJobType())
	if err != nil {
		return &models.GetJobResult{
			ErrorMessage: err.Error(),
			IsSuccessful: false,
		}, err
	}

	if job == nil {
		return &models.GetJobResult{
			Job:          nil,
			IsSuccessful: true,
		}, nil
	}

	protoJob := &models.Job{
		Id:        job.ID.String(),
		JobId:     job.JobID.String(),
		UserId:    job.UserID.String(),
		Status:    models.JobStatus(job.Status),
		Message:   &wrappers.StringValue{Value: job.GetMessage()},
		Context:   &wrappers.StringValue{Value: job.GetContext()},
		JobType:   job.JobType,
		CreatedAt: timestamppb.New(job.CreatedAt),
		FinishedAt: func() *timestamppb.Timestamp {
			if job.FinishedAt != nil {
				return timestamppb.New(*job.FinishedAt)
			}
			return nil
		}(),
		Progress: int32(job.Progress),
	}

	return &models.GetJobResult{
		Job:          protoJob,
		IsSuccessful: true,
	}, nil
}

func (j *JobRpcServer) GetJob(ctx context.Context, req *models.GetJobQuery) (*models.GetJobResult, error) {
	jobID, err := uuid.Parse(req.JobId)
	if err != nil {
		return &models.GetJobResult{IsSuccessful: false, ErrorMessage: err.Error()}, err
	}

	job, err := j.JobService.GetJobByJobID(jobID)
	if err != nil {
		return &models.GetJobResult{
			ErrorMessage: err.Error(),
			IsSuccessful: false,
		}, err
	}

	if job == nil {
		return &models.GetJobResult{
			Job:          nil,
			IsSuccessful: false,
		}, nil
	}

	protoJob := &models.Job{
		Id:        job.ID.String(),
		JobId:     job.JobID.String(),
		UserId:    job.UserID.String(),
		Status:    models.JobStatus(job.Status),
		Message:   &wrappers.StringValue{Value: job.GetMessage()},
		JobType:   job.JobType,
		CreatedAt: timestamppb.New(job.CreatedAt),
		FinishedAt: func() *timestamppb.Timestamp {
			if job.FinishedAt != nil {
				return timestamppb.New(*job.FinishedAt)
			}
			return nil
		}(),
		Progress: int32(job.Progress),
	}

	return &models.GetJobResult{
		Job:          protoJob,
		IsSuccessful: true,
	}, nil

}

func (j *JobRpcServer) CreateJob(ctx context.Context, req *models.CreateJobQuery) (*models.Result, error) {
	jobID, err := uuid.Parse(req.JobId)
	if err != nil {
		return &models.Result{IsSuccessful: false, ErrorMessage: err.Error()}, nil
	}

	userID, err := interceptors.GetUserIDFromContext(ctx)

	query := queries.CreateJobQuery{
		Message: &req.Message.Value,
		JobID:   jobID,
		JobType: req.JobType,
		Context: &req.Context.Value,
	}

	if err := j.JobService.CreateJob(query, userID); err != nil {
		return &models.Result{IsSuccessful: false, ErrorMessage: err.Error()}, nil
	}

	return &models.Result{IsSuccessful: true}, nil
}

func (j *JobRpcServer) StartJob(ctx context.Context, req *models.StartJobQuery) (*models.Result, error) {
	jobID, err := uuid.Parse(req.JobId)
	if err != nil {
		return &models.Result{IsSuccessful: false, ErrorMessage: err.Error()}, nil
	}

	userID, err := interceptors.GetUserIDFromContext(ctx)

	query := queries.StartJobQuery{
		JobID: jobID,
	}

	j.JobService.StartJob(query, userID)

	return &models.Result{IsSuccessful: true}, nil
}

func (j *JobRpcServer) CompleteJob(ctx context.Context, req *models.CompleteJobQuery) (*models.Result, error) {
	jobID, err := uuid.Parse(req.JobId)
	if err != nil {
		return &models.Result{IsSuccessful: false, ErrorMessage: err.Error()}, nil
	}

	userID, err := interceptors.GetUserIDFromContext(ctx)

	query := queries.CompleteJobQuery{
		JobID:    jobID,
		Message:  &req.Message.Value,
		IsFailed: req.IsFailed,
	}

	err = j.JobService.CompleteJob(query, userID)

	if err != nil {
		return &models.Result{IsSuccessful: false, ErrorMessage: err.Error()}, nil
	}

	return &models.Result{IsSuccessful: true}, nil
}
