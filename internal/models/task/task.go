package models

type (
	Task struct {
		ID          int    `json:"id"`
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
	}

	CreateTaskRequest struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
	}

	GetAllTaskResponse struct {
		Data       []Task     `json:"data"`
		Pagination Pagination `json:"pagination"`
	}

	Pagination struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}
)
