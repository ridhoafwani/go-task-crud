package services

import models "task-crud/internal/models/task"

func (s *Service) UpdateTask(id int, req models.CreateTaskRequest) (models.Task, error) {
	task := models.Task{
		Title:       req.Title,
		Description: req.Description,
	}

	updatedTask, err := s.taskRepository.Update(id, task)
	if err != nil {
		return models.Task{}, err
	}

	return updatedTask, nil
}
