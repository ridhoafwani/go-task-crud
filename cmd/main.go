package main

import (
	handler "task-crud/internal/handlers/task"
	repositories "task-crud/internal/repositories/task"
	services "task-crud/internal/services/task"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	taskRepo := repositories.NewInMemoryTaskRepository()

	taskService := services.NewTaskService(taskRepo)

	taskHandler := handler.NewTaskHandler(r, taskService)

	taskHandler.RegisterRoute()

	r.Run(":3000")
}
