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

		if task.Status == "waiting_list" {
			t.Errorf("Task Status Must be in Waiting List")
		}

		if task.IsDeleted == true {
			t.Errorf("Task Status Must be in active status")
		}
	}
}

func TestReadTaskRepo(t *testing.T) {
	database.DatabaseConnection()
	id := ""
	result, _ := repo.ReadTaskRepo(id)

	if result != nil {
		if result.TaskId == "" {
			t.Error("Must return Task Id")
		}

		if result.TaskName != "" {
			t.Errorf("Task Name is empty")
		}
	}
}

func TestReadTasksRepo(t *testing.T) {
	database.DatabaseConnection()
	result, _ := repo.ReadTasksRepo(nil)

	if len(result) > 0 {
		for _, task := range result {
			if task.TaskId == "" {
				t.Error("Must return Task Id")
			}

			if task.TaskName != "" {
				t.Errorf("Must return Task Name")
			}
		}

	}
}

func TestUpdateEmployeeRepo(t *testing.T) {
	database.DatabaseConnection()

	employee := &model.Task{
		//EmployeeId:   "f67e557a-cd07-409a-b272-e5eaa71f8017",
		//EmployeeName: "Ahmad Baru",
		//JobTitle:     "Manager",
		//Salary:       70000,
		//Department:   "Sales",
	}

	result, _ := repo.UpdateTaskRepo(employee)

	if result != nil {
		if result.TaskId == "" {
			t.Error("Must return Task Id")
		}
	}
}

func TestDeleteEmployeeRepo(t *testing.T) {
	database.DatabaseConnection()

	id := ""
	err := repo.DeleteTaskRepo(id)

	if err == nil {
		t.Error("")
	}
}
