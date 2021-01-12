package domain

import "time"

// Task data
type Task struct {
	ID          TaskID
	Title       TaskTitle
	TaskDate    TaskDate
	TaskOwner   TaskOwner
	TaskChecked TaskChecked
}

// New constructor
func New(id TaskID, title TaskTitle, date TaskDate, owner TaskOwner) Task {
	return Task{
		ID:        id,
		Title:     title,
		TaskDate:  date,
		TaskOwner: owner,
	}
}

// TaskID string id
type TaskID struct {
	value string
}

// Value of the task id
func (id TaskID) Value() string {
	return id.value
}

// Equals of the task id
func (id TaskID) Equals(other TaskID) bool {
	return id.value == other.Value()
}

// NewID ID Constructor
func NewID(value string) TaskID {
	return TaskID{
		value: value,
	}
}

// TaskTitle title
type TaskTitle struct {
	value string
}

// Value of title
func (title TaskTitle) Value() string {
	return title.value
}

// NewTitle Title Constructor
func NewTitle(value string) TaskTitle {
	return TaskTitle{
		value: value,
	}
}

// TaskDate creation date
type TaskDate struct {
	value time.Time
}

// Value time
func (date TaskDate) Value() time.Time {
	return date.value
}

// NewDate Creation date Constructor
func NewDate(value time.Time) TaskDate {
	return TaskDate{
		value: value,
	}
}

// TaskOwner owner
type TaskOwner struct {
	value string
}

// Value of owner
func (owner TaskOwner) Value() string {
	return owner.value
}

// NewOwner Owner Constructor
func NewOwner(value string) TaskOwner {
	return TaskOwner{
		value: value,
	}
}

// TaskChecked is checked
type TaskChecked struct {
	value bool
}

// Value of checked
func (checked TaskChecked) Value() bool {
	return checked.value
}

// NewChecked checked constructor
func NewChecked(value bool) TaskChecked {
	return TaskChecked{
		value: value,
	}
}
