package services

import models "task-crud/internal/models/task"

func (s *service) GetAllTasks(limit, offset int) (models.GetAllTaskResponse, error) {
	tasks, err := s.taskRepository.GetAll()
	if err != nil {
		return models.GetAllTaskResponse{}, err
	}

	start := offset
	end := offset + limit

	if start > len(tasks) {
		start = len(tasks)
	}
	if end > len(tasks) {
		end = len(tasks)
	}

	paginatedTasks := tasks[start:end]

	response := models.GetAllTaskResponse{
		Data: paginatedTasks,
		Pagination: models.Pagination{
			Limit:  limit,
			Offset: offset,
		},
	}

	return response, nil
}
