package repositories

import (
	"database/sql"
	"strings"

	"github.com/dictuantran/TaskPad/api/models"
)

// TaskRepository struct
type TaskRepository struct {
	db *sql.DB
}

// Init method
func (repo *TaskRepository) Init(db *sql.DB) {
	repo.db = db
}

// CreateTask method
func (repo *TaskRepository) CreateTask(task models.Task) (models.Task, error) {
	tags := strings.Join(task.Tags, ";")
	statement := `
    insert into tasks (userid, title, due, completed, effort, tags, notes)
    values ($1, $2, $3, $4, $5, $6, $7)
    returning id
  `
	var id int64
	err := repo.db.QueryRow(statement, task.UserID, task.Title, task.Due, task.Completed, task.Effort, tags, task.Notes).Scan(&id)
	if err != nil {
		return task, err
	}
	createdTask := task
	createdTask.ID = id
	return createdTask, nil
}
