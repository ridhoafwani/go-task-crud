package services

import (
	"context"
	models "task-crud/internal/models"
	"task-crud/utils/custom_err"
)

func (s *taskService) GetTaskByID(ctx context.Context, userID, id int64) (models.Task, error) {
	task, err := s.taskRepository.GetByID(ctx, id)
	if err != nil {
		return models.Task{}, err
	}

	if userID != task.UserID {
		return models.Task{}, custom_err.ErrUnauthorized
	}

	return task, err
}
