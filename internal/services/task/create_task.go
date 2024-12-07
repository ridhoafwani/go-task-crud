package services

import (
	"context"
	models "task-crud/internal/models"
)

func (s *taskService) CreateTask(ctx context.Context, userID int64, req models.CreateTaskRequest) (models.Task, error) {
	task := models.Task{
		Title:       req.Title,
		Description: req.Description,
	}

	createdTask, err := s.taskRepository.Create(ctx, userID, task)
	if err != nil {
		return models.Task{}, err
	}

	return createdTask, nil
}
