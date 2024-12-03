package services

import models "task-crud/internal/models/task"

func (s *service) GetTaskByID(id int) (models.Task, error) {
	task, err := s.taskRepository.GetByID(id)
	if err != nil {
		return models.Task{}, err
	}

	return task, err
}
