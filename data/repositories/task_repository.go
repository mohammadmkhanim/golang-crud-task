package repositories

import (
	"TaskCrud/data/models"
	"TaskCrud/utils"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

type TaskRepository struct {
	db *pgx.Conn
}

func NewTaskRepository(db *pgx.Conn) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(ctx context.Context, t *models.Task) error {
	query := `
	INSERT INTO tasks (id, title, description, status, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Exec(ctx, query,
		t.ID, t.Title, t.Description, t.Status, t.CreatedAt, t.UpdatedAt)

	if err != nil {
		utils.LogError("TaskRepository.Create", "failed to insert task {0}: {1}", t.ID, err)
	}

	return err
}

func (r *TaskRepository) GetAll(ctx context.Context) ([]models.Task, error) {
	rows, err := r.db.Query(ctx, "SELECT id, title, description, status, created_at, updated_at FROM tasks")
	if err != nil {
		utils.LogError("TaskRepository.GetAll", "failed to query tasks: {0}", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var t models.Task
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			utils.LogError("TaskRepository.GetAll", "failed to scan task row: {0}", err)
			return nil, err
		}
		tasks = append(tasks, t)
	}

	utils.LogSuccess("TaskRepository.GetAll", "fetched {0} tasks from database", len(tasks))

	return tasks, nil
}

func (r *TaskRepository) GetByID(ctx context.Context, id string) (*models.Task, error) {
	var t models.Task

	query := `
	SELECT id, title, description, status, created_at, updated_at
	FROM tasks WHERE id=$1
	`

	err := r.db.QueryRow(ctx, query, id).
		Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			utils.LogWarn("TaskRepository.GetByID", "task {0} not found", id)
		} else {
			utils.LogError("TaskRepository.GetByID", "failed to fetch task {0}: {1}", id, err)
		}
		return nil, err
	}

	return &t, nil
}

func (r *TaskRepository) Update(ctx context.Context, t *models.Task) error {
	query := `
	UPDATE tasks
	SET title=$1, description=$2, status=$3, updated_at=$4
	WHERE id=$5
	`

	_, err := r.db.Exec(ctx, query,
		t.Title, t.Description, t.Status, t.UpdatedAt, t.ID)

	if err != nil {
		utils.LogError("TaskRepository.Update", "failed to update task {0}: {1}", t.ID, err)
	}

	return err
}

func (r *TaskRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, "DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		utils.LogError("TaskRepository.Delete", "failed to delete task {0}: {1}", id, err)
	}
	return err
}
