package statistics_postgres_repository

import (
	"time"

	"github.com/jettyjunk/goland-todoapp/internal/core/domain"
)

type TaskModel struct {
	ID           int
	Version      int
	Title        string
	Description  *string
	Completed    bool
	CreatedAt    time.Time
	CompletedAt  *time.Time
	AuthorUserID int
}

func taskDomainFromModel(taskModel TaskModel) domain.Task {
	return domain.NewTask(
		taskModel.ID,
		taskModel.Version,
		taskModel.Title,
		taskModel.Description,
		taskModel.Completed,
		taskModel.CreatedAt,
		taskModel.CompletedAt,
		taskModel.AuthorUserID,
	)
}

func taskDomainsFromModels(tasksModels []TaskModel) []domain.Task {
	tasksDomains := make([]domain.Task, len(tasksModels))

	for index, task := range tasksModels {
		tasksDomains[index] = taskDomainFromModel(task)
	}

	return tasksDomains
}
