package services

import (
	"context"
	models "task-crud/internal/models"
	repositories "task-crud/internal/repositories/task"
)

type TaskService interface {
	CreateTask(ctx context.Context, userID int64, req models.CreateTaskRequest) (models.Task, error)
	GetAllTasks(ctx context.Context, userID int64, limit, offset int) (models.GetAllTaskResponse, error)
	GetTaskByID(ctx context.Context, userID, id int64) (models.Task, error)
	UpdateTask(ctx context.Context, userID, id int64, req models.UpdateTaskRequest) (models.Task, error)
	DeleteTask(ctx context.Context, userID, id int64) error
}

type taskService struct {
	taskRepository repositories.TaskRepository
}

func NewTaskService(repo repositories.TaskRepository) TaskService {
	return &taskService{
		taskRepository: repo,
	}
}
