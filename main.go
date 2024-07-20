package main

import (
	"fmt"

	"sagala-tech-test/database"
	emp "sagala-tech-test/internal/app/task/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting application ...")
	database.DatabaseConnection()

	r := gin.Default()
	r.GET("/task/:id", emp.ReadTask)
	r.POST("/tasks", emp.ReadTasks)
	r.POST("/task", emp.CreateTask)
	r.PUT("/task/:id", emp.UpdateTask)
	r.PUT("/task/:id/:status", emp.UpdateStatusTask)
	r.PUT("/task/:id/remove", emp.SoftDeleteTask)
	r.DELETE("/task/:id", emp.DeleteTask)

	r.Run(":5000")
}
