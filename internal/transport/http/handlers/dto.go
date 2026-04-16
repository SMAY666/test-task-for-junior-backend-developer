package handlers

import (
	"time"

	taskdomain "example.com/taskservice/internal/domain/task"
)

type taskMutationDTO struct {
	Title       string                `json:"title"`
	Description string                `json:"description"`
	IsRecurring bool                  `json:"is_recurring"`
	Recurring   *taskdomain.Recurring `json:"recurring"`
	Status      taskdomain.Status     `json:"status"`
	ExecutorIds []int32               `json:"executor_ids"`
}

type taskDTO struct {
	ID          int64                 `json:"id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	IsRecurring bool                  `json:"is_recurring"`
	Recurring   *taskdomain.Recurring `json:"recurring"`
	Status      taskdomain.Status     `json:"status"`
	ExecutorIds []int32               `json:"executor_ids"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
}

func newTaskDTO(task *taskdomain.Task) taskDTO {
	return taskDTO{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		IsRecurring: task.IsRecurring,
		Recurring:   task.Recurring,
		Status:      task.Status,
		ExecutorIds: task.ExecutorIds,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}
