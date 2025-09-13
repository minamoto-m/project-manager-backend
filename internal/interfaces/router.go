package interfaces

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/minamoto-m/project-manager-backend/internal/infrastructures/repositories"
	"github.com/minamoto-m/project-manager-backend/internal/interfaces/handlers"
	"github.com/minamoto-m/project-manager-backend/internal/usecases"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// CORS設定
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	taskRepo := repositories.NewTaskRepository(db)
	taskUsecase := usecases.NewTaskUsecase(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskUsecase)

	api := r.Group("/api/v1")
	{
		api.POST("/tasks", taskHandler.Create)
		api.GET("/tasks", taskHandler.GetAll)
		api.GET("/tasks/:id", taskHandler.GetByID)
	}

	return r
}
