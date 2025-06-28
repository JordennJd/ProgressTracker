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

	protoJob := &models.Job{
		Id:        job.ID.String(),
		JobId:     job.ID.String(),
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

func NewJobRpcServer(service *services.JobService) pb.JobServiceServer {
	return &JobRpcServer{
		JobService: service,
	}
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
