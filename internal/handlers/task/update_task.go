package handler

import (
	"net/http"
	"strconv"
	models "task-crud/internal/models/task"

	"github.com/gin-gonic/gin"
)

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid task ID",
		})
		return
	}

	var req models.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input request",
		})
		return
	}

	task, err := h.taskService.UpdateTask(id, req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, task)
}
