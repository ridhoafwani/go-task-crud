package repositories

import (
	"context"
	"database/sql"
	models "task-crud/internal/models"
)

func (repo *authRepository) GetUser(ctx context.Context, email, username string, id int64) (*models.User, error) {
	query := "SELECT id, email, username, password, created_at, updated_at FROM users WHERE email = $1 OR username = $2 OR id = $3"

	var user models.User
	err := repo.db.QueryRowContext(ctx, query, email, username, id).Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
