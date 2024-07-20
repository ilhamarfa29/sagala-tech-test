package task

import (
	"fmt"
	"testing"
	"time"

	"sagala-tech-test/database"
	model "sagala-tech-test/database/model"
	repo "sagala-tech-test/internal/app/task/repository"
)

func TestCreateTaskRepo(t *testing.T) {
	database.DatabaseConnection()

	layout := time.RFC3339
	str := "2024-07-20T10:10:30Z"
	timeParsed, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}

	task := &model.Task{
		TaskName:            "PR Matematika Pak Rusdi",
		Description:         "PR Matematika ini sulit, jadi perlu di breakdown satu-satu prosesnya dan dijelaskan",
		TaskDurationMinutes: 240,
		DueDate:             &timeParsed,
	}

	result, _ := repo.CreateTaskRepo(task)

	if result != nil {
		if result.TaskId == "" {
			t.Error("Must return Task Id")
		}

		if result.TaskName != task.TaskName {
			t.Errorf("Task Name is Different")
		}

		if result.Description != task.Description {
			t.Errorf("Description is Different")
		}

		if result.TaskDurationMinutes != task.TaskDurationMinutes {
			t.Errorf("Task Duration is Different")
		}

		if result.DueDate != task.DueDate {
			t.Errorf("Due Date is Different")
		}

		if result.Status != "waiting_list" {
			t.Errorf("Task Status Must be in Waiting List. Value %s", result.Status)
		}

		if result.IsDeleted {
			t.Errorf("Task Must be in active status. Value %t", result.IsDeleted)
		}
	}
}

func TestReadTaskRepo(t *testing.T) {
	database.DatabaseConnection()
	id := "ebbcda9b-d366-4496-8dff-68985957ba0e" // change Task Id if needed
	result, _ := repo.ReadTaskRepo(id)

	if result != nil {
		if result.TaskId == "" {
			t.Error("Must return Task Id")
		}

		if result.TaskName == "" {
			t.Errorf("Task Name is empty")
		}
	}
}

func TestReadTasksRepo(t *testing.T) {
	database.DatabaseConnection()

	filter := &model.Filter{
		Status:    "",
		IsDeleted: nil,
	}

	result, _ := repo.ReadTasksRepo(filter)

	if len(result) > 0 {
		for _, task := range result {
			if task.TaskId == "" {
				t.Error("Must return Task Id")
			}

			if task.TaskName == "" {
				t.Errorf("Must return Task Name")
			}
		}

	}
}

func TestUpdateTaskRepo(t *testing.T) {
	database.DatabaseConnection()

	id := "ebbcda9b-d366-4496-8dff-68985957ba0e" // change Task Id if needed

	layout := time.RFC3339
	str := "2024-07-20T10:10:30Z"
	timeParsed, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}

	task := &model.Task{
		TaskId:              id,
		TaskName:            "PR Matematika Pak Rusdi Updated",
		Description:         "PR Matematika ini sulit, jadi perlu di breakdown satu-satu prosesnya dan dijelaskan. Oh iya, dikumpulkan di kertas HVS",
		TaskDurationMinutes: 240,
		DueDate:             &timeParsed,
	}

	result, _ := repo.UpdateTaskRepo(task)

	if result != nil {
		if result.TaskId == "" {
			t.Error("Must return Task Id")
		}
	}
}

func TestUpdateStatusTask(t *testing.T) {
	database.DatabaseConnection()
	id := "ebbcda9b-d366-4496-8dff-68985957ba0e" // change Task Id if needed
	status := "done"                             // change Status if needed. Availabe=le value of status : waiting_list, in_progress, done
	resultData, _ := repo.ReadTaskRepo(id)
	if resultData == nil {
		t.Errorf("Task Id not found for soft delete process")
		return
	}

	resultData.Status = status

	result, err := repo.UpdateTaskRepo(resultData)
	if err != nil {
		t.Errorf("error : %s", err.Error())
		return
	}

	if result.TaskId != id {
		t.Errorf("Task Id is Different")
	}

	if result.Status != status {
		t.Errorf("Task Id Status not updated to %s", status)
	}
}

func TestSoftDeleteTask(t *testing.T) {
	database.DatabaseConnection()
	id := "ebbcda9b-d366-4496-8dff-68985957ba0e" // change Task Id if needed
	resultData, _ := repo.ReadTaskRepo(id)
	if resultData == nil {
		t.Errorf("Task Id not found for soft delete process")
		return
	}

	resultData.IsDeleted = true

	result, err := repo.UpdateTaskRepo(resultData)
	if err != nil {
		t.Errorf("error : %s", err.Error())
		return
	}

	if result.TaskId != id {
		t.Errorf("Task Id is Different")
	}

	if result.IsDeleted != true {
		t.Errorf("Task Id Status not updated to soft delete")
	}
}

func TestDeleteTaskRepo(t *testing.T) {
	database.DatabaseConnection()

	id := "87c96a80-3fbb-4c5d-802c-5e0987ee6881" // change Task Id if needed
	err := repo.DeleteTaskRepo(id)

	if err != nil {
		t.Error("")
	}
}
