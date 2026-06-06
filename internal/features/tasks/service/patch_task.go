package tasks_service

import (
	"context"
	"fmt"

	"github.com/jettyjunk/goland-todoapp/internal/core/domain"
)

func (s *TasksService) PatchTask(ctx context.Context, taskID int, patch domain.TaskPath) (domain.Task, error) {
	task, err := s.taskRepository.GetTask(ctx, taskID)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get taskID from repository: %w", err)
	}

	if err := task.ApplyPatch(patch); err != nil {
		return domain.Task{}, fmt.Errorf("apply task patch: %w", err)
	}

	patchedTask, err := s.taskRepository.PatchTask(ctx, taskID, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("patch task: %w", err)
	}

	return patchedTask, nil
}
