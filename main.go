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
	r.GET("/employees/:id", emp.ReadTask)
	r.GET("/employees", emp.ReadTasks)
	r.POST("/employees", emp.CreateTask)
	r.PUT("/employees/:id", emp.UpdateTask)
	r.DELETE("/employees/:id", emp.DeleteTask)

	r.Run(":5000")
}
