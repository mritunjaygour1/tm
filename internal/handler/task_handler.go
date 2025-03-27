package handler

import (
	"fmt"
	"net/http"
	"strings"
	"task-manager/internal/model"
	"task-manager/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gofrs/uuid"
)

type (
	ITaskHandler interface {
		GetTasks(*gin.Context)
		CreateTask(*gin.Context)
		GetTaskByID(*gin.Context)
		UpdateTaskByID(*gin.Context)
		DeleteTaskByID(*gin.Context)
	}

	TaskHandler struct {
		TaskService service.ITaskService
	}
)

const (
	ErrInvalidJSONBody       = "invalid JSON body"
	ErrInvalidRequestBody    = "invalid request body"
	ErrTaskNotFound          = "task not found"
	ErrTaskAlreadyCompleted  = "task already completed"
	ErrTaskAlreadyInProgress = "task already in progress"
	ErrTaskAlreadyPending    = "task already pending"
)

func NewTaskHandler(taskService service.ITaskService) *TaskHandler {
	return &TaskHandler{TaskService: taskService}
}

/*
	Handler functions
*/

func (h *TaskHandler) GetTasks(c *gin.Context) {
	ctx := c.Request.Context()

	// Fetch all tasks from the database
	tasks, err := h.TaskService.GetAllTasks(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &model.Response{Message: http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	ctx := c.Request.Context()
	var task model.Task

	// Bind the JSON body to the task model
	if err := c.ShouldBindJSON(&task); err != nil {
		errMsg := handleValidationError(err)
		c.JSON(http.StatusBadRequest, &model.Response{Messages: errMsg})
		return
	}

	task.ID, _ = uuid.NewV7()
	if err := h.TaskService.CreateTask(ctx, &task); err != nil {
		c.JSON(http.StatusInternalServerError, &model.Response{Message: http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	ctx := c.Request.Context()

	// Get the task ID from the URL
	id := c.Param("taskId")

	// Validate the task ID
	taskId, err := uuid.FromString(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &model.Response{Message: http.StatusText(http.StatusBadRequest)})
		return
	}

	// Fetch the task from the database
	task, err := h.TaskService.GetTaskByID(ctx, taskId)
	if err != nil {
		if strings.EqualFold(err.Error(), "record not found") {
			c.JSON(http.StatusNotFound, &model.Response{Message: ErrTaskNotFound})
			return
		}
		c.JSON(http.StatusInternalServerError, &model.Response{Message: http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) UpdateTaskByID(c *gin.Context) {
	ctx := c.Request.Context()

	// Get the task ID from the URL
	id := c.Param("taskId")

	// Validate the task ID
	taskId, err := uuid.FromString(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &model.Response{Message: http.StatusText(http.StatusBadRequest)})
		return
	}

	// Fetch the task from the database
	_, err = h.TaskService.GetTaskByID(ctx, taskId)
	if err != nil {
		if strings.EqualFold(err.Error(), "record not found") {
			c.JSON(http.StatusNotFound, &model.Response{Message: ErrTaskNotFound})
			return
		}
		c.JSON(http.StatusInternalServerError, &model.Response{Message: http.StatusText(http.StatusInternalServerError)})
		return
	}

	// Bind the JSON body to the task model
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		errMsg := handleValidationError(err)
		c.JSON(http.StatusBadRequest, &model.Response{Messages: errMsg})
		return
	}

	// Update the task in the database
	if err := h.TaskService.UpdateTask(ctx, taskId, &task); err != nil {
		c.JSON(http.StatusInternalServerError, &model.Response{Message: http.StatusText(http.StatusInternalServerError)})
		return
	}

	// Return the updated task
	c.JSON(http.StatusOK, &model.Response{Message: "Task updated successfully"})
}

func (h *TaskHandler) DeleteTaskByID(c *gin.Context) {
	ctx := c.Request.Context()

	// Get the task ID from the URL
	id := c.Param("taskId")

	// Validate the task ID
	taskId, err := uuid.FromString(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &model.Response{Message: http.StatusText(http.StatusBadRequest)})
		return
	}

	// Delete the task from the database
	if err := h.TaskService.DeleteTask(ctx, taskId); err != nil {
		c.JSON(http.StatusInternalServerError, &model.Response{Message: http.StatusText(http.StatusInternalServerError)})
		return
	}

	// Successfully delete the task
	c.JSON(http.StatusOK, &model.Response{Message: "Task deleted successfully"})
}

/*
	Suporting functions
*/

// handleValidationError customizes the error message when validation fails
func handleValidationError(err error) map[string]string {
	// Cast the error to a ValidationErrors type
	validationErrors, _ := err.(validator.ValidationErrors)

	// Create a map of custom error messages
	errorsMap := make(map[string]string)

	// Iterate through validation errors and generate custom error messages
	for _, e := range validationErrors {
		field := strings.ToLower(e.Field())
		switch e.Tag() {
		case "email":
			errorsMap[field] = "it must be a valid email address"
		case "max":
			errorsMap[field] = fmt.Sprintf("it must be at most %s characters long", e.Param())
		case "min":
			errorsMap[field] = fmt.Sprintf("it must be at least %s characters long", e.Param())
		case "name":
			errorsMap[field] = "it must adhere to valid naming conventions: no digits or special characters."
		case "oneof":
			errorsMap[field] = fmt.Sprintf("it must be one of the following [%s]", strings.Join(strings.Split(e.Param(), " "), ", "))
		case "phone":
			errorsMap[field] = "it must be a valid phone number"
		case "required":
			errorsMap[field] = "this is a required field"
		default:
			errorsMap[field] = "invalid value provided"
		}
	}

	return errorsMap
}
