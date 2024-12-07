package repositories

import (
	"context"
	"fmt"
	models "task-crud/internal/models"
)

func (repo *taskRepository) Create(ctx context.Context, userID int64, task models.Task) (models.Task, error) {
	query := `
        INSERT INTO tasks (title, description, user_id, created_at, updated_at)
        VALUES ($1, $2, $3, NOW(), NOW()) 
        RETURNING id, user_id, created_at, updated_at`

	err := repo.db.QueryRowContext(ctx, query, task.Title, task.Description, userID).Scan(&task.ID, &task.UserID, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		return task, fmt.Errorf("failed to create task: %w", err)
	}

	return task, nil
}
