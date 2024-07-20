package employee

// import (
// 	"testing"

// 	"sagala-tech-test/database"
// 	model "sagala-tech-test/database/model"
// 	repo "sagala-tech-test/internal/app/task/repository"
// )

// func TestCreateTaskRepo(t *testing.T) {
// 	database.DatabaseConnection()

// 	task := &model.Task{
// 		EmployeeName: "Ahmad",
// 		JobTitle:     "Manager",
// 		Salary:       70000,
// 		Department:   "Sales",
// 	}

// 	result, _ := repo.CreateTaskRepo(task)

// 	if result != nil {
// 		if result.EmployeeId == "" {
// 			t.Error("Must return Employee Id")
// 		}

// 		if result.EmployeeName != employee.EmployeeName {
// 			t.Errorf("Employee Name is Different")
// 		}

// 		if result.JobTitle != employee.JobTitle {
// 			t.Errorf("Job Title is Different")
// 		}

// 		if result.Salary != employee.Salary {
// 			t.Errorf("Salary is Different")
// 		}

// 		if result.Department != employee.Department {
// 			t.Errorf("Department is Different")
// 		}
// 	}
// }

// func TestReadEmployeeRepo(t *testing.T) {
// 	database.DatabaseConnection()
// 	id := ""
// 	result, _ := repo.ReadEmployeeRepo(id)

// 	if result != nil {
// 		if result.EmployeeId == "" {
// 			t.Error("Must return Employee Id")
// 		}

// 		if result.EmployeeName != "" {
// 			t.Errorf("Employee Name is empty")
// 		}

// 		if result.JobTitle != "" {
// 			t.Errorf("Job Title is empty")
// 		}

// 		if result.Salary != 0 {
// 			t.Errorf("Salary is empty")
// 		}

// 		if result.Department != "" {
// 			t.Errorf("Department is empty")
// 		}
// 	}
// }

// func TestReadEmployeesRepo(t *testing.T) {
// 	database.DatabaseConnection()
// 	result, _ := repo.ReadEmployeesRepo()

// 	if len(result) > 0 {
// 		for _, employee := range result {
// 			if employee.EmployeeId == "" {
// 				t.Error("Must return Employee Id")
// 			}
// 		}

// 	}
// }

// func TestUpdateEmployeeRepo(t *testing.T) {
// 	database.DatabaseConnection()

// 	employee := &model.Task{
// 		EmployeeId:   "f67e557a-cd07-409a-b272-e5eaa71f8017",
// 		EmployeeName: "Ahmad Baru",
// 		JobTitle:     "Manager",
// 		Salary:       70000,
// 		Department:   "Sales",
// 	}

// 	result, _ := repo.UpdateEmployeeRepo(employee)

// 	if result != nil {
// 		if result.EmployeeId == "" {
// 			t.Error("Must return Employee Id")
// 		}
// 	}
// }

// func TestDeleteEmployeeRepo(t *testing.T) {
// 	database.DatabaseConnection()

// 	id := ""
// 	err := repo.DeleteEmployeeRepo(id)

// 	if err == nil {
// 		t.Error("")
// 	}
// }
