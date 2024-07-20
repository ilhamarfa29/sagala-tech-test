package employee

import (
	"errors"
	"time"

	"sagala-tech-test/database"
	modelDb "sagala-tech-test/database/model"

	"github.com/google/uuid"
)

func CreateTaskRepo(task *modelDb.Task) (*modelDb.Task, error) {

	id := uuid.New()
	task.TaskId = id.String()
	if task.CreatedAt == nil {
		t := time.Now()
		task.CreatedAt = &t
	}

	res := database.DB.Create(task)
	if res.RowsAffected == 0 {
		return nil, errors.New("error creating a task")
	}

	return task, nil
}

func ReadTaskRepo(id string) (*modelDb.Task, error) {

	var task modelDb.Task

	res := database.DB.Where("task_id = ?", id).Find(&task)
	if res.RowsAffected == 0 {
		return nil, errors.New("task not found")
	}

	return &task, nil
}

func ReadTasksRepo() ([]modelDb.Task, error) {

	var task []modelDb.Task
	res := database.DB.Find(&task)
	if res.Error != nil {
		return nil, errors.New("task not found")
	}

	return task, nil
}

func UpdateTaskRepo(task *modelDb.Task) (*modelDb.Task, error) {
	var updateTask modelDb.Task
	res := database.DB.Model(&updateTask).Where("task_id = ?", task.TaskId).Updates(task)

	if res.RowsAffected == 0 {
		return nil, errors.New("task not updated")
	}

	return task, nil
}

func DeleteTaskRepo(id string) error {
	var employee modelDb.Task
	res := database.DB.Where("task_id = ?", id).Find(&employee)
	if res.RowsAffected == 0 {
		return errors.New("task not found")
	}
	database.DB.Delete(&employee, "task_id = ?", id)

	return nil
}
