package handler

import (
	services "task-crud/internal/services/task"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	*gin.Engine
	taskService services.TaskService
}

func NewTaskHandler(api *gin.Engine, s services.TaskService) *TaskHandler {
	return &TaskHandler{
		Engine:      api,
		taskService: s,
	}
}

func (h *TaskHandler) RegisterRoute() {
	router := h.Group("api/v1/tasks")

	router.POST("/", h.CreateTask)
	router.GET("/", h.GetAllTasks)
	router.GET("/:id", h.GetTaskByID)
	router.PATCH("/:id", h.UpdateTask)
	router.DELETE("/:id", h.DeleteTask)
}
