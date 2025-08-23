package api

import (
	"github.com/gin-gonic/gin"
	"github.com/minamoto-m/project-manager-backend/api/handlers"
)

func SetupRoutes(r *gin.Engine) {
	// OpenAPIスキーマに基づくプロジェクトハンドラー
	projectHandler := &handlers.ProjectHandler{}

	api := r.Group("/api/v1")
	{
		// プロジェクト関連のルート
		api.GET("/projects", projectHandler.GetProjects)
		api.POST("/projects", projectHandler.PostProjects)
		// api.GET("/projects/:projectId", projectHandler.GetProjectsProjectId)
		// api.PATCH("/projects/:projectId", projectHandler.PatchProjectsProjectId)
		// api.DELETE("/projects/:projectId", projectHandler.DeleteProjectsProjectId)
	}
}
