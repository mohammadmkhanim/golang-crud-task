package repositories

import (
	"TaskCrud/data/models"
	"context"

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

	return err
}

func (r *TaskRepository) GetAll(ctx context.Context) ([]models.Task, error) {
	rows, err := r.db.Query(ctx, "SELECT id, title, description, status, created_at, updated_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var t models.Task
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

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

	return err
}

func (r *TaskRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, "DELETE FROM tasks WHERE id=$1", id)
	return err
}
