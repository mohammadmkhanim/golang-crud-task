package services

import (
	"context"
	"time"

	"TaskCrud/data/models"
	"TaskCrud/data/repositories"

	"github.com/google/uuid"
)

type TaskService struct {
	repo *repositories.TaskRepository
}

func NewTaskService(repo *repositories.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(ctx context.Context, t *models.Task) error {
	t.ID = uuid.New().String()
	t.Status = models.Todo
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()

	return s.repo.Create(ctx, t)
}

func (s *TaskService) GetAll(ctx context.Context) ([]models.Task, error) {
	return s.repo.GetAll(ctx)
}

func (s *TaskService) GetByID(ctx context.Context, id string) (*models.Task, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *TaskService) UpdateTask(ctx context.Context, t *models.Task) error {
	t.UpdatedAt = time.Now()
	return s.repo.Update(ctx, t)
}

func (s *TaskService) DeleteTask(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
