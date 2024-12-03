package repositories

import "errors"

func (repo *InMemoryTaskRepository) Delete(id int) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	_, exists := repo.tasks[id]
	if !exists {
		return errors.New("task not found")
	}
	delete(repo.tasks, id)
	return nil
}
