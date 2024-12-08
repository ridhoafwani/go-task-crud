package repositories

import (
	"context"
	"database/sql"
	models "task-crud/internal/models"
	"time"
)

func (repo *authRepository) GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*models.RefreshToken, error) {
	query := "SELECT id, user_id, refresh_token, expired_at, created_at, updated_at FROM refresh_tokens WHERE user_id = $1 AND expired_at >= $2"

	var refreshToken models.RefreshToken
	err := repo.db.QueryRowContext(ctx, query, userID, now).Scan(&refreshToken.ID, &refreshToken.UserID, &refreshToken.RefreshToken, &refreshToken.ExpiredAt, &refreshToken.CreatedAt, &refreshToken.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &refreshToken, nil
}
