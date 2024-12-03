package repositories

import models "task-crud/internal/models/task"

func (repo *InMemoryTaskRepository) Create(task models.Task) (models.Task, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	task.ID = repo.nextID
	repo.tasks[repo.nextID] = task
	repo.nextID++
	return task, nil
}
