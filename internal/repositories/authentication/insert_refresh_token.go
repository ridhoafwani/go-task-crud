package repositories

import (
	"context"
	models "task-crud/internal/models"
)

func (repo *authRepository) InsertRefreshToken(ctx context.Context, model models.RefreshToken) error {
	query := `
        INSERT INTO refresh_tokens (user_id, refresh_token, expired_at, created_at, updated_at)
        VALUES ($1, $2, $3, NOW(), NOW()) `

	_, err := repo.db.ExecContext(ctx, query, model.UserID, model.RefreshToken, model.ExpiredAt)

	if err != nil {
		return err
	}

	return nil
}
