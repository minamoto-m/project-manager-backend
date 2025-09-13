package repositories

import (
	"github.com/minamoto-m/project-manager-backend/internal/domains"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Save(task *domains.Task) error {
	return r.db.Create(task).Error
}

func (r *TaskRepository) FindAll() ([]domains.Task, error) {
	var tasks []domains.Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) FindByID(id uint) (*domains.Task, error) {
	var task domains.Task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}
