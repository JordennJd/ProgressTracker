package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go/types"
	"net/http"
	"progress-tracker/internal/middlewares"
	"progress-tracker/internal/queries"
	"progress-tracker/internal/responses"
	"progress-tracker/internal/services"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type JobHandler struct {
	service         services.JobService
	progressService services.ProgressService
}

func NewJobHandler(service services.JobService, progressService services.ProgressService) *JobHandler {
	return &JobHandler{service: service, progressService: progressService}
}

func (h *JobHandler) CreateJob(w http.ResponseWriter, r *http.Request) {
	var input queries.CreateJobQuery
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, "invalid query: "+err.Error(), http.StatusBadRequest)
		return
	}

	userIDStr := r.Context().Value(middlewares.UserIDKey).(string)
	uuidValue, err := uuid.Parse(userIDStr)
	if err != nil {
		respondWithError(w, "Invalid UUID format", http.StatusUnauthorized)
		return
	}

	err = h.service.CreateJob(input, uuidValue)
	if err != nil {
		respondWithError(w, "job create error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithSuccessful(w, &input.JobID)
}

func (h *JobHandler) StartJob(w http.ResponseWriter, r *http.Request) {
	var input queries.StartJobQuery
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, "invalid query: "+err.Error(), http.StatusBadRequest)
		return
	}
	if input.JobID == uuid.Nil {
		respondWithError(w, "invalid job id", http.StatusBadRequest)
		return
	}
	userIDStr := r.Context().Value(middlewares.UserIDKey).(string)
	userId, err := uuid.Parse(userIDStr)
	if err != nil {
		respondWithError(w, "Invalid UUID format", http.StatusUnauthorized)
		return
	}

	err = h.service.StartJob(input, userId)
	if err != nil {
		respondWithError(w, "job start error: "+err.Error(), http.StatusBadRequest)
		return
	}

	respondWithSuccessful(w, &input.JobID)
}

func (h *JobHandler) GetJobByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		respondWithError(w, "invalid UUID", http.StatusBadRequest)
		return
	}

	job, err := h.service.GetJobByID(id)
	if err != nil {
		respondWithError(w, err.Error(), http.StatusBadRequest)
	}

	respondWithSuccessful(w, &job)
}

func (h *JobHandler) GetJobByJobID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["job_id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		respondWithError(w, "invalid UUID", http.StatusBadRequest)
		return
	}

	job, err := h.service.GetJobByJobID(id)
	if err != nil {
		respondWithError(w, err.Error(), http.StatusBadRequest)
	}

	respondWithSuccessful(w, &job)
}

func (h *JobHandler) GetAllJob(w http.ResponseWriter, r *http.Request) {
	jobs, err := h.service.GetAll()

	if err != nil {
		respondWithError(w, "internal error", http.StatusBadRequest)
	}

	respondWithSuccessful(w, &jobs)
}

func (h *JobHandler) SetJobProgress(w http.ResponseWriter, r *http.Request) {
	var input queries.SetProgressQuery
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, "invalid query: "+err.Error(), http.StatusBadRequest)
		return
	}
	if input.JobID == uuid.Nil {
		respondWithError(w, "invalid job id", http.StatusBadRequest)
		return
	}

	isJobExist := h.service.IsJobExists(input.JobID)
	if !isJobExist {
		respondWithError(w, "job does not exist", http.StatusBadRequest)
	}

	err := h.progressService.SetProgress(input)
	if err != nil {
		respondWithError(w, "job start error: "+err.Error(), http.StatusBadRequest)
		return
	}

	respondWithSuccessful(w, &input.JobID)
}

func (h *JobHandler) GetProgress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["job_id"]
	fmt.Println(idStr)

	jobID, err := uuid.Parse(idStr)
	if err != nil {
		respondWithError(w, "invalid UUID", http.StatusBadRequest)
		return
	}
	progress, err := h.progressService.GetProgress(jobID)

	if err != nil {
		respondWithError(w, "get progress error"+err.Error(), http.StatusBadRequest)
		return
	}

	respondWithSuccessful(w, &progress)
}

func respondWithError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(responses.Result[types.Nil]{
		ErrorMessage: message,
		IsSuccessful: false,
	})
}

func respondWithSuccessful[T any](w http.ResponseWriter, data *T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses.Result[T]{
		Data:         data,
		IsSuccessful: true,
	})
}
