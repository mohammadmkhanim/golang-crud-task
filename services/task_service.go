package services

import (
	"TaskCrud/data/models"
	"TaskCrud/data/repositories"
	"TaskCrud/utils"
	"context"
)

type TaskService struct {
	taskRepository *repositories.TaskRepository
}

func NewTaskService(repo *repositories.TaskRepository) *TaskService {
	return &TaskService{taskRepository: repo}
}

func (s *TaskService) CreateTask(ctx context.Context, t *models.Task) error {
	var nowTime = utils.NowUTC()

	t.ID = utils.NewDbId()
	t.Status = models.Todo
	t.CreatedAt = nowTime
	t.UpdatedAt = nowTime

	return s.taskRepository.Create(ctx, t)
}

func (s *TaskService) GetAll(ctx context.Context, status *models.TaskStatus, order models.SortOrder) ([]models.Task, error) {
	return s.taskRepository.GetAll(ctx, status, order)
}

func (s *TaskService) GetByID(ctx context.Context, id string) (*models.Task, error) {
	return s.taskRepository.GetByID(ctx, id)
}

func (s *TaskService) UpdateTask(ctx context.Context, t *models.Task, input *models.Task) error {
	t.Title = input.Title
	t.Description = input.Description
	t.Status = input.Status
	t.UpdatedAt = utils.NowUTC()

	return s.taskRepository.Update(ctx, t)
}

func (s *TaskService) DeleteTask(ctx context.Context, id string) error {
	return s.taskRepository.Delete(ctx, id, utils.NowUTC())
}
