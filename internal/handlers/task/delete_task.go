package handler

import (
	"net/http"
	"strconv"
	"task-crud/utils/custom_err"

	"github.com/gin-gonic/gin"
)

// @Summary Delete a task by ID
// @Description Delete a task by ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} models.MessageResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /tasks/{id} [delete]
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	userID := c.GetInt64("userID")
	err = h.taskService.DeleteTask(c, userID, id)
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

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
