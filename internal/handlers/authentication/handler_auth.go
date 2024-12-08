package handler

import (
	"task-crud/internal/middleware"
	services "task-crud/internal/services/authentication"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	*gin.Engine
	authService services.AuthService
}

func NewAuthHandler(api *gin.Engine, s services.AuthService) *AuthHandler {
	return &AuthHandler{
		Engine:      api,
		authService: s,
	}
}

func (h *AuthHandler) RegisterRoute() {
	router := h.Group("api/v1/auth")

	router.POST("/signup", h.SignUp)
	router.POST("/signin", h.SignIn)

	// Refresh token
	refreshRouter := h.Group("api/v1/auth")
	refreshRouter.Use(middleware.AuthRefreshMiddleware())
	refreshRouter.POST("/refresh", h.RefreshToken)
}
