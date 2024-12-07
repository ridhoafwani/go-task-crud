package main

import (
	"task-crud/internal/config"
	"task-crud/internal/database"

	authHandler "task-crud/internal/handlers/authentication"
	taskHandler "task-crud/internal/handlers/task"

	authRepository "task-crud/internal/repositories/authentication"
	taskRepository "task-crud/internal/repositories/task"

	_ "task-crud/docs"
	authService "task-crud/internal/services/authentication"
	taskService "task-crud/internal/services/task"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Task CRUD API
// @version 1.0
// @description This is a simple CRUD API for managing tasks
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	r := gin.Default()

	config.LoadConfig()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	dbService := database.New()

	// Get the DB connection instance
	db := dbService.GetDB()

	taskRepo := taskRepository.NewTaskRepository(db)
	authRepo := authRepository.NewAuthRepository(db)

	taskService := taskService.NewTaskService(taskRepo)
	authService := authService.NewAuthService(authRepo)

	taskHandler := taskHandler.NewTaskHandler(r, taskService)
	authHandler := authHandler.NewAuthHandler(r, authService)

	taskHandler.RegisterRoute()
	authHandler.RegisterRoute()

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":" + config.AppConfig.Port)
}
