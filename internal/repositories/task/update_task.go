package repositories

import (
	"errors"
	models "task-crud/internal/models/task"
)

func (repo *InMemoryTaskRepository) Update(id int, task models.Task) (models.Task, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	_, exists := repo.tasks[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}
	task.ID = id
	repo.tasks[id] = task
	return task, nil
}
