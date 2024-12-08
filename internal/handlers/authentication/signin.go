package handler

import (
	"net/http"
	models "task-crud/internal/models"
	"task-crud/utils/custom_err"

	"github.com/gin-gonic/gin"
)

// @Summary Sign in a user
// @Description Sign in a user with the input payload
// @Tags auth
// @Accept json
// @Produce json
// @Param request body models.SignInRequest true "Sign in request"
// @Success 200 {object} models.SignInResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /auth/signin [post]
func (h *AuthHandler) SignIn(c *gin.Context) {
	ctx := c.Request.Context()

	var request models.SignInRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	accessToken, refreshToken, err := h.authService.SignIn(ctx, request)
	if err != nil {
		if err == custom_err.ErrCredentialInvalid {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	response := models.SignInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	c.JSON(http.StatusOK, response)
}
