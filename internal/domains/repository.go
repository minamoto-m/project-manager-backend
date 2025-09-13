package domains

type TaskRepository interface {
	Save(task *Task) error
	FindByID(id uint) (*Task, error)
	FindAll() ([]Task, error)
}
