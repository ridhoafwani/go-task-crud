package main

import (
	handler "task-crud/internal/handlers/task"
	repositories "task-crud/internal/repositories/task"
	services "task-crud/internal/services/task"

	_ "task-crud/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Task CRUD API
// @version 1.0
// @description This is a simple CRUD API for managing tasks
// @host localhost:3000
// @BasePath /api/v1

func main() {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	taskRepo := repositories.NewInMemoryTaskRepository()

	taskService := services.NewTaskService(taskRepo)

	taskHandler := handler.NewTaskHandler(r, taskService)

	taskHandler.RegisterRoute()

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":3000")
}
