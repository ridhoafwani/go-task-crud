package services

import models "task-crud/internal/models/task"

func (s *service) CreateTask(req models.CreateTaskRequest) (models.Task, error) {
	task := models.Task{
		Title:       req.Title,
		Description: req.Description,
	}

	createdTask, err := s.taskRepository.Create(task)
	if err != nil {
		return models.Task{}, err
	}

	return createdTask, nil
}
