package controllers

import (
	"database/sql"

	"github.com/dictuantran/TaskPad/api/repositories"
)

// TaskController struct
type TaskController struct {
	taskRepository *repositories.TaskRepository
}

// Init method
func (c *TaskController) Init(db *sql.DB) {
	c.taskRepository = &repositories.TaskRepository{}
	c.taskRepository.Init(db)
}
