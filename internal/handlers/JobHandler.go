package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"progress-tracker/internal/services"

	"progress-tracker/internal/models"
)

type JobHandler struct {
	service services.JobService
}

func NewJobHandler(service services.JobService) *JobHandler {
	return &JobHandler{service: service}
}

func (h *JobHandler) CreateJob(w http.ResponseWriter, r *http.Request) {
	var input models.Job
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Неверный запрос: "+err.Error(), http.StatusBadRequest)
		return
	}

	err := h.service.SaveJob(&input)
	if err != nil {
		http.Error(w, "Ошибка сохранения в БД: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(input)
}

func (h *JobHandler) GetJobByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid UUID", http.StatusBadRequest)
		return
	}

	job, err := h.service.GetJobByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}
