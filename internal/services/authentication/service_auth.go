package services

import (
	"context"
	models "task-crud/internal/models"
	repositories "task-crud/internal/repositories/authentication"
)

type AuthService interface {
	SignUp(ctx context.Context, req models.SignUpRequest) error
	SignIn(ctx context.Context, req models.SignInRequest) (string, string, error)
	ValidateRefreshToken(ctx context.Context, userID int64, request models.RefreshTokenRequest) (string, error)
}

type authService struct {
	repo repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) AuthService {
	return &authService{
		repo: repo,
	}
}
