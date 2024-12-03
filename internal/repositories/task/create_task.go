package repositories

import models "task-crud/internal/models/task"

func (repo *InMemoryTaskRepository) Create(task models.Task) (models.Task, error) {
	task.ID = repo.nextID
	repo.tasks[repo.nextID] = task
	repo.nextID++
	return task, nil
}
