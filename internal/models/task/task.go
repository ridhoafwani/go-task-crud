package models

// Task represents a task in the system
// @Description Task information
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// CreateTaskRequest represents the request body for creating a task
// @Description Create task request
type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type (
	GetAllTaskResponse struct {
		Data       []Task     `json:"data"`
		Pagination Pagination `json:"pagination"`
	}

	Pagination struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type MessageResponse struct {
	Message string `json:"message"`
}
