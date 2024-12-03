package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get a task by ID
// @Description Get a task by ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} models.Task
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /tasks/{id} [get]
func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := h.taskService.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}
