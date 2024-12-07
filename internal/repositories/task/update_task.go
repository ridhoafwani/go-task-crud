package repositories

import (
	"context"
	"database/sql"
	"fmt"
	models "task-crud/internal/models"
	"task-crud/utils/custom_err"
)

func (repo *taskRepository) Update(ctx context.Context, id int64, task models.Task) (models.Task, error) {
	query := "UPDATE tasks SET title = $1, description = $2, updated_at = NOW() WHERE id = $3 RETURNING updated_at"

	err := repo.db.QueryRowContext(ctx, query, task.Title, task.Description, id).Scan(&task.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return task, custom_err.ErrNotFound
		}
		return task, fmt.Errorf("failed to update task: %w", err)
	}

	return task, nil
}
