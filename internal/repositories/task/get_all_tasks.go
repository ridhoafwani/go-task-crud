package repositories

import models "task-crud/internal/models/task"

func (repo *InMemoryTaskRepository) GetAll() ([]models.Task, error) {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()

	var taskList []models.Task
	for _, task := range repo.tasks {
		taskList = append(taskList, task)
	}
	return taskList, nil
}
