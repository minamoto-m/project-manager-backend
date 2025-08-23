package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProjectHandler implements the ServerInterface
type ProjectHandler struct{}

// モックデータ
var mockProjects = []Project{
	{
		Id:          stringPtr("1"),
		Name:        stringPtr("プロジェクトA"),
		Description: stringPtr("これはテストプロジェクトAです"),
		Color:       stringPtr("#FF5733"),
	},
	{
		Id:          stringPtr("2"),
		Name:        stringPtr("プロジェクトB"),
		Description: stringPtr("これはテストプロジェクトBです"),
		Color:       stringPtr("#33FF57"),
	},
}

func stringPtr(s string) *string {
	return &s
}

// GetProjects - プロジェクト一覧取得
func (h *ProjectHandler) GetProjects(c *gin.Context) {
	response := ProjectListResponse{
		Projects: &mockProjects,
	}
	c.JSON(http.StatusOK, response)
}

// PostProjects - プロジェクト作成
func (h *ProjectHandler) PostProjects(c *gin.Context) {
	var request ProjectCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		errorResponse := ErrorResponse{
			Message: stringPtr("リクエストの形式が正しくありません"),
			Code:    stringPtr("VALIDATION_ERROR"),
		}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// バリデーション
	if request.Name == nil || *request.Name == "" {
		errorResponse := ErrorResponse{
			Message: stringPtr("プロジェクト名は必須です"),
			Code:    stringPtr("VALIDATION_ERROR"),
		}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// 新しいプロジェクトを作成
	newProject := Project{
		Id:          stringPtr(strconv.Itoa(len(mockProjects) + 1)),
		Name:        request.Name,
		Description: request.Description,
		Color:       request.Color,
	}

	response := ProjectResponse{
		Project: &newProject,
	}
	c.JSON(http.StatusCreated, response)
}

// GetProjectsProjectId - プロジェクト詳細取得
func (h *ProjectHandler) GetProjectsProjectId(c *gin.Context, projectId string) {
	// プロジェクトを検索
	for _, project := range mockProjects {
		if project.Id != nil && *project.Id == projectId {
			response := ProjectResponse{
				Project: &project,
			}
			c.JSON(http.StatusOK, response)
			return
		}
	}

	// プロジェクトが見つからない場合
	errorResponse := ErrorResponse{
		Message: stringPtr("プロジェクトが見つかりません"),
		Code:    stringPtr("NOT_FOUND"),
	}
	c.JSON(http.StatusNotFound, errorResponse)
}

// PatchProjectsProjectId - プロジェクト更新
func (h *ProjectHandler) PatchProjectsProjectId(c *gin.Context, projectId string) {
	var request ProjectUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		errorResponse := ErrorResponse{
			Message: stringPtr("リクエストの形式が正しくありません"),
			Code:    stringPtr("VALIDATION_ERROR"),
		}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	// プロジェクトを検索して更新
	for i, project := range mockProjects {
		if project.Id != nil && *project.Id == projectId {
			if request.Name != nil {
				mockProjects[i].Name = request.Name
			}
			if request.Description != nil {
				mockProjects[i].Description = request.Description
			}
			if request.Color != nil {
				mockProjects[i].Color = request.Color
			}

			response := ProjectResponse{
				Project: &mockProjects[i],
			}
			c.JSON(http.StatusOK, response)
			return
		}
	}

	// プロジェクトが見つからない場合
	errorResponse := ErrorResponse{
		Message: stringPtr("プロジェクトが見つかりません"),
		Code:    stringPtr("NOT_FOUND"),
	}
	c.JSON(http.StatusNotFound, errorResponse)
}

// DeleteProjectsProjectId - プロジェクト削除
func (h *ProjectHandler) DeleteProjectsProjectId(c *gin.Context, projectId string) {
	// プロジェクトを検索して削除
	for i, project := range mockProjects {
		if project.Id != nil && *project.Id == projectId {
			// スライスから削除
			mockProjects = append(mockProjects[:i], mockProjects[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}

	// プロジェクトが見つからない場合
	errorResponse := ErrorResponse{
		Message: stringPtr("プロジェクトが見つかりません"),
		Code:    stringPtr("NOT_FOUND"),
	}
	c.JSON(http.StatusNotFound, errorResponse)
}
