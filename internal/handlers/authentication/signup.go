package handler

import (
	"net/http"
	models "task-crud/internal/models"
	"task-crud/utils/custom_err"

	"github.com/gin-gonic/gin"
)

// @Summary Sign up a new user
// @Description Sign up a new user with the input payload
// @Tags auth
// @Accept json
// @Produce json
// @Param request body models.SignUpRequest true "Sign up request"
// @Success 201 {object} models.MessageResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /auth/signup [post]
func (h *AuthHandler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()

	var request models.SignUpRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.authService.SignUp(ctx, request)
	if err != nil {
		if err == custom_err.ErrUserNameOrEmailAlreadyExist {
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

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}
