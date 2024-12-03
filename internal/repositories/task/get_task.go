package repositories

import (
	"errors"
	models "task-crud/internal/models/task"
)

func (repo *InMemoryTaskRepository) GetByID(id int) (models.Task, error) {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()

	task, exists := repo.tasks[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}
	return task, nil
}
