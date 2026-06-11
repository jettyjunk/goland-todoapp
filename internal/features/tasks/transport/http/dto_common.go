package tasks_transport_http

import (
	"time"

	"github.com/jettyjunk/goland-todoapp/internal/core/domain"
)

type TaskDTOResponse struct {
	ID           int        `json:"id" example:"1"`
	Version      int        `json:"version" example:"1"`
	Title        string     `json:"title" example:"Сходить на тренировку"`
	Description  *string    `json:"description" example:"Сделать..."`
	Completed    bool       `json:"completed" example:"false"`
	CreatedAt    time.Time  `json:"created_at" example:"2026-02-26T12:00:00Z"`
	CompletedAt  *time.Time `json:"completed_at" example:"null"`
	AuthorUserID int        `json:"author_user_id" example:"2"`
}

func taskDTOFromDomain(task domain.Task) TaskDTOResponse {
	return TaskDTOResponse{
		ID:           task.ID,
		Version:      task.Version,
		Title:        task.Title,
		Description:  task.Description,
		Completed:    task.Completed,
		CreatedAt:    task.CreatedAt,
		CompletedAt:  task.CompletedAt,
		AuthorUserID: task.AuthorUserID,
	}
}

func tasksDTOFromDomains(tasks []domain.Task) []TaskDTOResponse {
	tasksDTO := make([]TaskDTOResponse, len(tasks))

	for index, task := range tasks {
		tasksDTO[index] = taskDTOFromDomain(task)
	}

	return tasksDTO
}
