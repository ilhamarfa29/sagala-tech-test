package database

import "time"

type Task struct {
	TaskId              string     `json:"task_id"`
	TaskName            string     `json:"task_name"`
	Description         string     `json:"description"`
	Status              string     `json:"status"`
	IsDeleted           bool       `json:"is_deleted"`
	TaskDurationMinutes int        `json:"task_duration_minutes"`
	DueDate             *time.Time `json:"due_date"`
	CreatedAt           *time.Time `json:"created_at"`
}
