package repository

import (
	"lotest/internal/model"
	"sync"
)

type Repo struct {
	tasks map[string]model.Task
	mu    sync.RWMutex
}

func NewRepo() *Repo {
	return &Repo{
		tasks: make(map[string]model.Task),
	}
}

func (r *Repo) Add(task model.Task) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tasks[task.ID] = task
}

func (r *Repo) FindByID(id string) (model.Task, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	task, exists := r.tasks[id]
	return task, exists
}

func (r *Repo) FindAll() []model.Task {
	r.mu.RLock()
	defer r.mu.RUnlock()
	tasks := make([]model.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (r *Repo) FindByStatus(status string) []model.Task {
	r.mu.RLock()
	defer r.mu.RUnlock()
	tasks := []model.Task{}
	for _, task := range r.tasks {
		if task.Status == status {
			tasks = append(tasks, task)
		}
	}
	return tasks
}
