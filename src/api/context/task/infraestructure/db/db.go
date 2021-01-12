package db

import (
	"fmt"

	"github.com/rlgino/monda-todo/src/api/context/task/domain"
)

// InMemoryDB implementation
type InMemoryDB struct {
	tasks []domain.Task
}

// SaveTask persist the task
func (db *InMemoryDB) SaveTask(t domain.Task) error {
	db.tasks = append(db.tasks, t)
	len := len(db.tasks)
	fmt.Println(fmt.Sprintf("New insert: %d", len))
	return nil
}

// UpdateTask update the id task
func (db *InMemoryDB) UpdateTask(t domain.Task) error {
	for i, el := range db.tasks {
		if t.ID.Equals(el.ID) {
			db.tasks[i] = t
			return nil
		}
	}

	return fmt.Errorf("The id wasn't found")
}

// FindTask into memory DB
func (db *InMemoryDB) FindTask(id domain.TaskID) (*domain.Task, error) {
	for _, el := range db.tasks {
		if el.ID.Equals(id) {
			return &el, nil
		}
	}

	return nil, nil
}

// FindTasks find all tasks into memory DB
func (db *InMemoryDB) FindTasks(owner domain.TaskOwner) ([]domain.Task, error) {
	result := []domain.Task{}
	for _, el := range db.tasks {
		if el.TaskOwner.Value() == owner.Value() {
			result = append(result, el)
		}
	}
	return result, nil
}
