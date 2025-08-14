package handler

import (
	"encoding/json"
	"lotest/internal/logger"
	"lotest/internal/model"
	"lotest/internal/service"
	"net/http"
)

type TaskHandler struct {
	service *service.TaskService
	logger  *logger.Logger
}

func NewTaskHandler(s *service.TaskService, l *logger.Logger) *TaskHandler {
	return &TaskHandler{service: s, logger: l}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	var tasks []model.Task

	if status != "" {
		tasks = h.service.GetTasksByStatus(status)
	} else {
		tasks = h.service.GetAllTasks()
		h.logger.Log("GET_ALL", "All tasks retrieved")
	}

	respondJSON(w, tasks)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	task, exists := h.service.GetTask(id)
	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		h.logger.Log("ERROR", "Task not found: "+id)
		return
	}

	h.logger.Log("GET", "Task retrieved: "+id)
	respondJSON(w, task)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		h.logger.Log("ERROR", "Invalid create request: "+err.Error())
		return
	}

	task := h.service.CreateTask(req.Title, req.Description)
	h.logger.Log("CREATE", "Task created: "+task.ID)
	w.WriteHeader(http.StatusCreated)
	respondJSON(w, task)
}

func respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
