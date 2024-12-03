package repositories

import (
	"sync"
	models "task-crud/internal/models/task"
)

type InMemoryTaskRepository struct {
	tasks  map[int]models.Task
	nextID int
	mutex  sync.RWMutex
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks:  make(map[int]models.Task),
		nextID: 0,
	}
}

type TaskRepository interface {
	Create(task models.Task) (models.Task, error)
	GetAll() ([]models.Task, error)
	GetByID(id int) (models.Task, error)
	Update(id int, task models.Task) (models.Task, error)
	// Delete(id int) error
}
