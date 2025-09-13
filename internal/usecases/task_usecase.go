package usecases

import (
	"time"

	"github.com/minamoto-m/project-manager-backend/internal/domains"
	"github.com/minamoto-m/project-manager-backend/internal/dto/request"
)

type TaskUsecase struct {
	taskRepo domains.TaskRepository
}

func NewTaskUsecase(repo domains.TaskRepository) *TaskUsecase {
	return &TaskUsecase{taskRepo: repo}
}

func (u *TaskUsecase) Create(req *request.TaskCreateRequest) (*domains.Task, error) {
	// DTOからドメインエンティティを生成
	task := &domains.Task{
		ProjectID:   req.ProjectID,
		Title:       req.Title,
		Description: req.Description,
		Priority:    uint(req.Priority),
		DueDate:     parseDueDate(req.DueDate),
		CreaterID:   1, // 仮の値（後で認証システムから取得）
		AssigneeID:  1, // 仮の値（後で認証システムから取得）
		CreatedAt:   time.Now(),
	}

	if err := u.taskRepo.Save(task); err != nil {
		return nil, err
	}

	return task, nil
}

// parseDueDate 文字列の日付をtime.Timeに変換
func parseDueDate(dueDateStr string) time.Time {
	if dueDateStr == "" {
		return time.Time{}
	}

	parsed, err := time.Parse("2006-01-02", dueDateStr)
	if err != nil {
		return time.Time{}
	}

	return parsed
}

func (u *TaskUsecase) GetAll() ([]domains.Task, error) {
	return u.taskRepo.FindAll()
}

func (u *TaskUsecase) GetByID(id uint) (*domains.Task, error) {
	return u.taskRepo.FindByID(id)
}
