package services

import (
	models "task-crud/internal/models/task"
	repositories "task-crud/internal/repositories/task"
)

type TaskService interface {
	CreateTask(req models.CreateTaskRequest) (models.Task, error)
	// GetAllTasks(limit, offset int) (models.GetAllTaskResponse, error)
	// GetTaskByID(id int) (models.Task, error)
	// UpdateTask(id int, req models.CreateTaskRequest) (models.Task, error)
	// DeleteTask(id int) error
}

type service struct {
	taskRepository repositories.TaskRepository
}

func NewTaskService(repo repositories.TaskRepository) *service {
	return &service{
		taskRepository: repo,
	}
}
