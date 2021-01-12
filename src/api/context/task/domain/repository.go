package domain

// TaskRepository interface to save
type TaskRepository interface {
	SaveTask(Task) error
	UpdateTask(Task) error
	FindTask(TaskID) (*Task, error)
	FindTasks(TaskOwner) ([]Task, error)
}
