package repositories

import (
	"context"
	models "task-crud/internal/models"
)

func (repo *authRepository) CreateUser(ctx context.Context, model models.User) error {
	query := `
        INSERT INTO users (email, username, password, created_at, updated_at)
        VALUES ($1, $2, $3, NOW(), NOW()) `

	_, err := repo.db.ExecContext(ctx, query, model.Email, model.Username, model.Password)

	if err != nil {
		return err
	}

	return nil
}
