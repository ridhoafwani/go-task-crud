package handler

// @title Task CRUD API
// @version 1.0
// @description This is a simple CRUD API for managing tasks
// @host localhost:3000
// @BasePath /api/v1

import (
	"task-crud/internal/middleware"
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

	router.Use(middleware.AuthMiddleware())

	router.POST("/", h.CreateTask)
	router.GET("/", h.GetAllTasks)
	router.GET("/:id", h.GetTaskByID)
	router.PATCH("/:id", h.UpdateTask)
	router.DELETE("/:id", h.DeleteTask)
}
