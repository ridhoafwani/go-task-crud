package models

import "time"

type (
	SignUpRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Username string `json:"username" binding:"required,min=3,max=100"`
		Password string `json:"password" binding:"required,min=8"`
	}

	SignInRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	RefreshTokenRequest struct {
		Token string `json:"token"`
	}
)

type (
	SignInResponse struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	}

	RefreshResponse struct {
		AccessToken string `json:"accessToken"`
	}
)

type (
	User struct {
		ID        int64     `db:"id"`
		Email     string    `db:"email"`
		Username  string    `db:"username"`
		Password  string    `db:"password"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}

	RefreshToken struct {
		ID           int64     `db:"id"`
		UserID       int64     `db:"user_id"`
		RefreshToken string    `db:"refresh_token"`
		ExpiredAt    time.Time `db:"expired_at"`
		CreatedAt    time.Time `db:"created_at"`
		UpdatedAt    time.Time `db:"updated_at"`
	}
)
