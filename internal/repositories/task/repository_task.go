package repositories

import (
	"context"
	"database/sql"
	models "task-crud/internal/models"
)

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db: db}
}

type TaskRepository interface {
	Create(ctx context.Context, userID int64, task models.Task) (models.Task, error)
	GetAll(ctx context.Context, userID int64, limit, offset int) ([]models.Task, error)
	GetByID(ctx context.Context, id int64) (models.Task, error)
	Update(ctx context.Context, id int64, task models.Task) (models.Task, error)
	Delete(ctx context.Context, id int64) error
}
