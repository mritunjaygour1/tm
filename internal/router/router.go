package router

import (
	"task-manager/internal/handler"
	"task-manager/internal/middleware"
	"task-manager/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, TaskService service.ITaskService) {
	healthzHandler := handler.NewHealthzHandler()
	taskHandler := handler.NewTaskHandler(TaskService)

	// Healthz endpoint
	activity := router.Group("/activity")
	activity.GET("/healthz", healthzHandler.GetHealthz) // Get Health status

	// Task endpoints
	tasks := router.Group("/tasks")
	tasks.GET("/", taskHandler.GetTasks) // Get All Tasks

	tasks.Use(middleware.AuthMiddleware)                 // Auth Middleware added
	tasks.POST("/", taskHandler.CreateTask)              // Create Task
	tasks.GET("/:taskId", taskHandler.GetTaskByID)       // Get Task by ID
	tasks.PUT("/:taskId", taskHandler.UpdateTaskByID)    // Update Task by ID
	tasks.DELETE("/:taskId", taskHandler.DeleteTaskByID) // Delete Task by ID
}
