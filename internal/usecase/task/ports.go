package task

import (
	"context"

	taskdomain "example.com/taskservice/internal/domain/task"
)

type Repository interface {
	Create(ctx context.Context, task *taskdomain.Task) (*taskdomain.Task, error)
	GetByID(ctx context.Context, id string) (*taskdomain.Task, error)
	Update(ctx context.Context, task *taskdomain.Task) (*taskdomain.Task, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]taskdomain.Task, error)
}

type Usecase interface {
	Create(ctx context.Context, input CreateInput) (*taskdomain.Task, error)
	GetByID(ctx context.Context, id string) (*taskdomain.Task, error)
	Update(ctx context.Context, id string, input UpdateInput) (*taskdomain.Task, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]taskdomain.Task, error)
}

type CreateInput struct {
	Title       string
	Description string
	IsRecurring bool
	Recurring   *taskdomain.Recurring
	Status      taskdomain.Status
	ExecutorIds []string
}

type UpdateInput struct {
	Title       string
	Description string
	IsRecurring bool
	Recurring   *taskdomain.Recurring
	Status      taskdomain.Status
	ExecutorIds []string
}
