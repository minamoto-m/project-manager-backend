package domains

import "time"

type Task struct {
	ID          uint
	ProjectID   uint
	Title       string
	Description string
	Priority    uint
	DueDate     time.Time
	CreaterID   uint
	AssigneeID  uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
