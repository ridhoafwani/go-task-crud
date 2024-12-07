package services

import (
	"context"
	models "task-crud/internal/models"
	"task-crud/utils/custom_err"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s *authService) SignUp(ctx context.Context, req models.SignUpRequest) error {
	user, err := s.repo.GetUser(ctx, req.Email, req.Username, 0)
	if err != nil {
		return err
	}

	if user != nil {
		return custom_err.ErrUserNameOrEmailAlreadyExist
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	model := models.User{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(pass),
		CreatedAt: now,
		UpdatedAt: now,
	}
	return s.repo.CreateUser(ctx, model)
}
