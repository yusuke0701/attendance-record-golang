package domain

import (
	"context"
	"time"
)

type TaskID string

func (t TaskID) String() string {
	return string(t)
}

type (
	TaskRepository interface {
		Create(context.Context, Task) (Task, error)
		FindAll(context.Context) ([]Task, error)
		WithTransaction(context.Context, func(context.Context) error) error
	}

	Task struct {
		id TaskID

		created_at time.Time
		updated_at time.Time
	}
)

func NewTask(
	ID TaskID,
	created_at time.Time,
	updated_at time.Time,
) Task {
	return Task{
		id:         ID,
		created_at: created_at,
		updated_at: updated_at,
	}
}

func (t Task) ID() TaskID {
	return t.id
}

func (t Task) CreatedAt() time.Time {
	return t.created_at
}

func (t Task) UpdatedAt() time.Time {
	return t.updated_at
}
