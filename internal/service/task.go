package service

import (
	"lotest/internal/model"
	"lotest/internal/repository"
	"strconv"
	"time"
)

type TaskService struct {
	repo *repository.Repo
}

func NewTaskService(repo *repository.Repo) *TaskService {
	return &TaskService{repo}
}

func (s *TaskService) CreateTask(title, description string) model.Task {
	task := model.Task{
		ID:          generateID(),
		Title:       title,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
	}
	s.repo.Add(task)
	return task
}

func (s *TaskService) GetTask(id string) (model.Task, bool) {
	return s.repo.FindByID(id)
}

func (s *TaskService) GetAllTasks() []model.Task {
	return s.repo.FindAll()
}

func (s *TaskService) GetTasksByStatus(status string) []model.Task {
	return s.repo.FindByStatus(status)
}

func generateID() string {
	return strconv.Itoa(time.Now().Nanosecond())
}
