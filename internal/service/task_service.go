package service

import (
	"context"
	"task-manager/internal/model"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

type (
	ITaskService interface {
		CreateTask(context.Context, *model.Task) error
		GetAllTasks(context.Context) ([]model.Task, error)
		GetTaskByID(context.Context, uuid.UUID) (*model.Task, error)
		UpdateTask(context.Context, uuid.UUID, *model.Task) error
		DeleteTask(context.Context, uuid.UUID) error
	}

	TaskService struct {
		DB *gorm.DB
	}
)

func NewTaskService(db *gorm.DB) ITaskService {
	return &TaskService{DB: db}
}

func (s *TaskService) CreateTask(ctx context.Context, task *model.Task) error {
	return s.DB.Create(task).Error
}

func (s *TaskService) GetAllTasks(ctx context.Context) ([]model.Task, error) {
	var tasks []model.Task
	err := s.DB.Find(&tasks).Error
	return tasks, err
}

func (s *TaskService) GetTaskByID(ctx context.Context, id uuid.UUID) (*model.Task, error) {
	var task model.Task
	err := s.DB.First(&task, "id = ?", id).Error
	return &task, err
}

func (s *TaskService) UpdateTask(ctx context.Context, id uuid.UUID, task *model.Task) error {
	return s.DB.Model(&model.Task{}).Where("id = ?", id).Updates(task).Error
}

func (s *TaskService) DeleteTask(ctx context.Context, id uuid.UUID) error {
	return s.DB.Delete(&model.Task{}, "id = ?", id).Error
}
