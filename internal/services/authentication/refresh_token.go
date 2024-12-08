package services

import (
	"context"
	"errors"
	"task-crud/internal/config"
	models "task-crud/internal/models"
	"task-crud/utils/jwt"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *authService) ValidateRefreshToken(ctx context.Context, userID int64, request models.RefreshTokenRequest) (string, error) {
	log.Info().Msg("Validating refresh token")
	existingRefreshToken, err := s.repo.GetRefreshToken(ctx, userID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("error get refresh token from database")
		return "", err
	}

	if existingRefreshToken == nil {
		return "", errors.New("refresh token has expired")
	}

	if existingRefreshToken.RefreshToken != request.Token {
		return "", errors.New("refresh token is invalid")
	}

	user, err := s.repo.GetUser(ctx, "", "", userID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}
	if user == nil {
		return "", errors.New("user not exist")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, config.AppConfig.JWTSecretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}
