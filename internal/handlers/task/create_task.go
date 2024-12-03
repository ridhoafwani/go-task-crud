package handler

import (
	"errors"
	"net/http"
	models "task-crud/internal/models/task"

	"github.com/gin-gonic/gin"
)

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req models.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("invalid input request").Error(),
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
