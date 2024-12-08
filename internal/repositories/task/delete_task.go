package repositories

import (
	"context"
	"fmt"
	"task-crud/utils/custom_err"
)

func (repo *taskRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM tasks WHERE id = $1"
	result, err := repo.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return custom_err.ErrNotFound
	}

	return nil
}
