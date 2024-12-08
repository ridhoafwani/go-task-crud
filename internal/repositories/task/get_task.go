package repositories

import (
	"context"
	"database/sql"
	"fmt"
	models "task-crud/internal/models"
	"task-crud/utils/custom_err"
)

func (repo *taskRepository) GetByID(ctx context.Context, id int64) (models.Task, error) {
	query := "SELECT id, title, description, user_id, created_at, updated_at FROM tasks WHERE id = $1"

	var task models.Task
	err := repo.db.QueryRowContext(ctx, query, id).Scan(&task.ID, &task.Title, &task.Description, &task.UserID, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return task, custom_err.ErrNotFound
		}
		return task, fmt.Errorf("failed to get task: %w", err)
	}

	return task, nil
}
