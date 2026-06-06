package tasks_service

import (
	"context"

	"github.com/jettyjunk/goland-todoapp/internal/core/domain"
)

type TasksService struct {
	taskRepository TaskRepository
}

type TaskRepository interface {
	CreateTask(
		ctx context.Context,
		task domain.Task,
	) (domain.Task, error)

	GetTasks(
		ctx context.Context,
		userID *int,
		limit *int,
		offset *int,
	) ([]domain.Task, error)

	GetTask(
		ctx context.Context,
		taskID int,
	) (domain.Task, error)

	DeleteTask(
		ctx context.Context,
		taskID int,
	) error

	PatchTask(
		ctx context.Context,
		taskID int,
		patch domain.Task,
	) (domain.Task, error)
}

func NewTasksService(taskRepository TaskRepository) *TasksService {
	return &TasksService{
		taskRepository: taskRepository,
	}
}
