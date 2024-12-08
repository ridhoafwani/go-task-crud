package services

import (
	"context"
	models "task-crud/internal/models"
)

func (s *taskService) UpdateTask(ctx context.Context, userID, id int64, req models.UpdateTaskRequest) (models.Task, error) {

	existingTask, err := s.GetTaskByID(ctx, userID, id)
	if err != nil {
		return existingTask, err
	}

	existingTask.Title = req.Title
	existingTask.Description = req.Description

	return s.taskRepository.Update(ctx, id, existingTask)

}
