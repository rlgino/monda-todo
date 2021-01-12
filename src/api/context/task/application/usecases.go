package application

import (
	"time"

	"github.com/rlgino/monda-todo/src/api/context/task/domain"
)

// TaskUseCases use cases of different tasks
type TaskUseCases interface {
	CreateTask(ID string, title string, owner string, creationDate time.Time) error
	DeleteTask(ID string) error
	UpdateTaskTitle(ID string, title string) error
	CheckTask(ID string) error
	FindTask(ID string) (domain.Task, error)
	FindTasks(owner string) ([]domain.Task, error)
}

// New constructor
func New(repo domain.TaskRepository) TaskUseCases {
	return &useCases{
		repo: repo,
	}
}

type useCases struct {
	repo domain.TaskRepository
}

func (uc *useCases) CreateTask(ID string, title string, owner string, creationDate time.Time) error {
	id := domain.NewID(ID)
	titleVO := domain.NewTitle(title)
	dateVO := domain.NewDate(creationDate)
	ownerVO := domain.NewOwner(owner)

	task := domain.New(id, titleVO, dateVO, ownerVO)
	uc.repo.SaveTask(task)
	return nil
}

func (uc *useCases) DeleteTask(ID string) error {
	return nil
}

func (uc *useCases) UpdateTaskTitle(ID string, title string) error {
	id := domain.NewID(ID)

	task, err := uc.repo.FindTask(id)
	if err != nil {
		return err
	}
	if task == nil {
		return domain.NotFoundError{}
	}

	titleVO := domain.NewTitle(title)

	task.Title = titleVO

	uc.repo.UpdateTask(*task)
	return nil
}

func (uc *useCases) CheckTask(ID string) error {
	id := domain.NewID(ID)

	task, err := uc.repo.FindTask(id)
	if err != nil {
		return err
	}
	if task == nil {
		return domain.NotFoundError{}
	}

	checked := domain.NewChecked(!task.TaskChecked.Value())

	task.TaskChecked = checked

	uc.repo.UpdateTask(*task)
	return nil
}

func (uc *useCases) FindTask(ID string) (domain.Task, error) {
	id := domain.NewID(ID)

	task, err := uc.repo.FindTask(id)
	if err != nil {
		return domain.Task{}, err
	}
	if task == nil {
		return domain.Task{}, domain.NotFoundError{}
	}

	return *task, nil
}

func (uc *useCases) FindTasks(owner string) ([]domain.Task, error) {
	taskOwner := domain.NewOwner(owner)
	tasks, err := uc.repo.FindTasks(taskOwner)
	if err != nil {
		return []domain.Task{}, err
	}

	return tasks, nil
}
