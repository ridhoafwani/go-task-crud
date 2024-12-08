package services

import (
	"context"
	models "task-crud/internal/models"
)

func (s *taskService) GetAllTasks(ctx context.Context, userID int64, limit, offset int) (models.GetAllTaskResponse, error) {
	tasks, err := s.taskRepository.GetAll(ctx, userID, limit, offset)
	if err != nil {
		return models.GetAllTaskResponse{}, err
	}

	response := models.GetAllTaskResponse{
		Data: tasks,
		Pagination: models.Pagination{
			Limit:  limit,
			Offset: offset,
		},
	}

	return response, nil
}
