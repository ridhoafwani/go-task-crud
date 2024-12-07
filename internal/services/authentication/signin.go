package services

import (
	"context"
	"errors"
	"time"

	"task-crud/internal/config"
	models "task-crud/internal/models"
	"task-crud/utils/custom_err"
	"task-crud/utils/jwt"
	token_util "task-crud/utils/token"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *authService) SignIn(ctx context.Context, req models.SignInRequest) (string, string, error) {
	user, err := s.repo.GetUser(ctx, req.Email, "", 0)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", "", err
	}
	if user == nil {
		return "", "", custom_err.ErrCredentialInvalid
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", "", custom_err.ErrCredentialInvalid
	}

	token, err := jwt.CreateToken(user.ID, user.Username, config.AppConfig.JWTSecretKey)
	if err != nil {
		return "", "", err
	}

	existingRefreshToken, err := s.repo.GetRefreshToken(ctx, user.ID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("error get latest refresh token from database")
		return "", "", err
	}

	if existingRefreshToken != nil {
		return token, existingRefreshToken.RefreshToken, nil
	}

	refreshToken := token_util.GenerateRefreshToken()
	if refreshToken == "" {
		return token, "", errors.New("failed to generate refresh token")
	}

	err = s.repo.InsertRefreshToken(ctx, models.RefreshToken{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiredAt:    time.Now().Add(10 * 24 * time.Hour),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		log.Error().Err(err).Msg("error inserting refresh token to database")
		return token, refreshToken, err
	}

	return token, refreshToken, nil
}
