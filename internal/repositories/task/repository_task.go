package repositories

import model "task-crud/internal/models/task"

type InMemoryTaskRepository struct {
	tasks  map[int]model.Task
	nextID int
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks:  make(map[int]model.Task),
		nextID: 1,
	}
}

type TaskRepository interface {
	Create(task model.Task) (model.Task, error)
	// GetAll() ([]model.Task, error)
	// GetByID(id int) (model.Task, error)
	// Update(id int, task model.Task) (model.Task, error)
	// Delete(id int) error
}
