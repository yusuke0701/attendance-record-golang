package usecase

import (
	"context"
	"time"

	"attendance-record/domain"
)

type (
	CreateTaskUsecase interface {
		Execute(context.Context, CreateTaskInput) (CreateTaskOutput, error)
	}

	CreateTaskInput struct {
		Name string `json:"name"`
	}

	CreateTaskPresenter interface {
		Output(domain.Task) CreateTaskOutput
	}

	CreateTaskOutput struct {
		ID   string `json:"id"`
		Name string `json:"name"`

		CreatedAt string `json:"created_at"`
		UpdateAt  string `json:"update_at"`
	}

	createTaskInteractor struct {
		taskRepo  domain.TaskRepository
		presenter CreateTaskPresenter
	}
)

func NewCreateTaskInteractor(
	taskRepo domain.TaskRepository,
	presenter CreateTaskPresenter,
) CreateTaskUsecase {
	return createTaskInteractor{
		taskRepo:  taskRepo,
		presenter: presenter,
	}
}

func (t createTaskInteractor) Execute(ctx context.Context, input CreateTaskInput) (CreateTaskOutput, error) {
	var (
		task domain.Task
		err  error
	)

	// 実際にはトランザクションは不要だが、サンプルとして用意した
	err = t.taskRepo.WithTransaction(ctx, func(ctx context.Context) error {
		task = domain.NewTask(
			domain.TaskID(domain.NewUUID()),
			time.Now(),
			time.Now(),
		)

		task, err = t.taskRepo.Create(ctx, task)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return t.presenter.Output(domain.Task{}), err
	}

	return t.presenter.Output(task), nil
}
