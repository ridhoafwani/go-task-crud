package repositories

import (
	"context"
	"fmt"
	models "task-crud/internal/models"
)

func (repo *taskRepository) GetAll(ctx context.Context, userID int64, limit, offset int) ([]models.Task, error) {
	query := "SELECT id, title, description, user_id, created_at, updated_at FROM tasks WHERE user_id = $1 LIMIT $2 OFFSET $3"

	rows, err := repo.db.QueryContext(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tasks: %w", err)
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.UserID, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed during row iteration: %w", err)
	}

	return tasks, nil
}
