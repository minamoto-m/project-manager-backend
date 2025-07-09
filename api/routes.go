package api

import (
	"github.com/gin-gonic/gin"
	"github.com/minamoto-m/project-manager-backend/api/handlers"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/projects", handlers.ListProjects)
		api.POST("/projects", handlers.CreateProject)
		api.PUT("/projects/:id", handlers.UpdateProject)
		api.DELETE("/projects/:id", handlers.DeleteProject)
		api.GET("/tasks", handlers.ListTasks)
		api.POST("/tasks", handlers.CreateTask)
		api.PUT("/tasks/:id", handlers.UpdateTask)
		api.DELETE("/tasks/:id", handlers.DeleteTask)
	}
}
