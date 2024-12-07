package repositories

import (
	"context"
	"database/sql"
	models "task-crud/internal/models"
	"time"
)

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{db: db}
}

type AuthRepository interface {
	GetUser(ctx context.Context, email, username string, userID int64) (*models.User, error)
	CreateUser(ctx context.Context, model models.User) error
	GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*models.RefreshToken, error)
	InsertRefreshToken(ctx context.Context, model models.RefreshToken) error
}
