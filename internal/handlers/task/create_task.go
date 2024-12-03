package handler

import (
	"net/http"
	models "task-crud/internal/models/task"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new task
// @Description Create a new task with the input payload
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body models.CreateTaskRequest true "Create task request"
// @Success 201 {object} models.Task
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /tasks [post]
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req models.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input request",
		})
		return
	}

	task, err := h.taskService.CreateTask(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, task)
}
