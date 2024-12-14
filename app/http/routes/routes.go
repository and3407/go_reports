package routes

import (
	"github.com/and3407/go_reports/app/http/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(ginEngine *gin.Engine) {
	reportRoutes(ginEngine)
	taskRoutes(ginEngine)
}

func reportRoutes(ginEngine *gin.Engine) {
	controller := controllers.ReportController{}

	reportRoutes := ginEngine.Group("/reports")
	reportRoutes.POST("/", controller.AddReport)
}

func taskRoutes(ginEngine *gin.Engine) {
	controller := controllers.TaskController{}

	taskRoutes := ginEngine.Group("/tasks")
	taskRoutes.POST("/", controller.AddTask)
}