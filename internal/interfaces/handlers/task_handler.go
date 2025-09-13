package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minamoto-m/project-manager-backend/internal/dto/request"
	"github.com/minamoto-m/project-manager-backend/internal/dto/response"
	"github.com/minamoto-m/project-manager-backend/internal/usecases"
)

type TaskHandler struct {
	taskUsecase *usecases.TaskUsecase
}

func NewTaskHandler(taskUsecase *usecases.TaskUsecase) *TaskHandler {
	return &TaskHandler{taskUsecase: taskUsecase}
}

func (h *TaskHandler) Create(c *gin.Context) {
	var req request.TaskCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// DTOをUsecaseに渡して、ドメイン生成を実行。
	// 戻り値として作成済みのドメインエンティティを受け取る。
	task, err := h.taskUsecase.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// ドメイン → レスポンスDTOへマッピングし、Locationヘッダを付けて201で返却
	response := response.TaskResponse{
		ID:          task.ID,
		ProjectID:   task.ProjectID,
		Title:       task.Title,
		Description: task.Description,
		Priority:    task.Priority,
		DueDate:     task.DueDate.Format(time.RFC3339),
		CreatorID:   task.CreaterID,
		AssigneeID:  task.AssigneeID,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
	c.Header("Location", fmt.Sprintf("/api/v1/tasks/%d", task.ID))
	c.JSON(http.StatusCreated, response)
}

// GetAll タスク一覧を取得
func (h *TaskHandler) GetAll(c *gin.Context) {
	tasks, err := h.taskUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetByID 指定されたIDのタスクを取得
func (h *TaskHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	task, err := h.taskUsecase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}
