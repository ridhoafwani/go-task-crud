package handler

import (
	"net/http"
	"strconv"
	models "task-crud/internal/models"
	"task-crud/utils/custom_err"

	"github.com/gin-gonic/gin"
)

// @Summary Update a task by ID
// @Description Update a task by ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param request body models.CreateTaskRequest true "Task information"
// @Success 200 {object} models.Task
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /tasks/{id} [patch]
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid task ID",
		})
		return
	}

	var req models.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input request",
		})
		return
	}

	userID := c.GetInt64("userID")
	task, err := h.taskService.UpdateTask(c, userID, id, req)
	if err != nil {
		if err == custom_err.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else if err == custom_err.ErrUnauthorized {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, task)
}
