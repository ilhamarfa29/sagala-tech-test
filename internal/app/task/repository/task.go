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

	res := database.DB.Table(modelDb.TableNameTask).Create(task)
	if res.RowsAffected == 0 {
		return nil, errors.New("error creating a task")
	}

	return task, nil
}

func ReadTaskRepo(id string) (*modelDb.Task, error) {

	var task modelDb.Task

	res := database.DB.Table(modelDb.TableNameTask).Where("task_id = ?", id).Find(&task)
	if res.RowsAffected == 0 {
		return nil, errors.New("task not found")
	}

	return &task, nil
}

func ReadTasksRepo(filter *modelDb.Filter) ([]modelDb.Task, error) {

	var task []modelDb.Task
	if filter.Status == "" && filter.IsDeleted == nil {
		res := database.DB.Table(modelDb.TableNameTask).Order("created_at desc").Find(&task)
		if res.Error != nil {
			return nil, errors.New("task not found")
		}
	} else if filter.Status != "" && filter.IsDeleted != nil {
		res := database.DB.Table(modelDb.TableNameTask).Where(map[string]interface{}{"status": filter.Status, "is_deleted": filter.IsDeleted}).
			Order("created_at desc").Find(&task)
		if res.Error != nil {
			return nil, errors.New("task not found")
		}
	} else if filter.Status != "" {
		res := database.DB.Table(modelDb.TableNameTask).Where("status = ?", filter.Status).Order("created_at desc").Find(&task)
		if res.Error != nil {
			return nil, errors.New("task not found")
		}
	} else if filter.IsDeleted != nil {
		res := database.DB.Table(modelDb.TableNameTask).Where("is_deleted = ?", filter.IsDeleted).Order("created_at desc").Find(&task)
		if res.Error != nil {
			return nil, errors.New("task not found")
		}
	}

	return task, nil
}

func UpdateTaskRepo(task *modelDb.Task) (*modelDb.Task, error) {
	var updateTask modelDb.Task
	res := database.DB.Table(modelDb.TableNameTask).Model(&updateTask).Where("task_id = ?", task.TaskId).Updates(task)

	if res.RowsAffected == 0 {
		return nil, errors.New("task not updated")
	}

	return task, nil
}

func DeleteTaskRepo(id string) error {
	var task modelDb.Task
	res := database.DB.Table(modelDb.TableNameTask).Where("task_id = ?", id).Find(&task)
	if res.RowsAffected == 0 {
		return errors.New("task not found")
	}
	database.DB.Table(modelDb.TableNameTask).Delete(&task, "task_id = ?", id)

	return nil
}
