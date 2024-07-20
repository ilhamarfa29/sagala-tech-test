package employee

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	modelDb "sagala-tech-test/database/model"
	repo "sagala-tech-test/internal/app/task/repository"
	constants "sagala-tech-test/internal/constant"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var (
	domainNameTask = "task"
)

func initLog() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)

	return log
}

func CreateTask(c *gin.Context) {
	log := initLog()
	log.Info(fmt.Printf(constants.CreateProcess, domainNameTask))
	defer log.Info(fmt.Printf(constants.CreateProcessDone, domainNameTask))

	var task *modelDb.Task
	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		log.Error(err.Error())
		return
	}

	id := uuid.New()
	task.TaskId = id.String()

	result, err := repo.CreateTaskRepo(task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a task",
		})
		log.Error("error creating a task")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": result,
	})

	log.Info(fmt.Printf(constants.CreateProcessSuccess, domainNameTask))

	return
}

func ReadTask(c *gin.Context) {
	log := initLog()
	log.Info(fmt.Printf(constants.ReadByIdProcess, domainNameTask))
	defer log.Info(fmt.Printf(constants.ReadByIdProcessDone, domainNameTask))

	id := c.Param("id")
	result, err := repo.ReadTaskRepo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "task not found",
		})

		log.Error("task not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": result,
	})
	log.Info(fmt.Printf(constants.ReadByIdProcessSuccess, domainNameTask))

	return
}

func ReadTasks(c *gin.Context) {
	log := initLog()
	log.Info(fmt.Printf(constants.ReadProcess, domainNameTask))
	defer log.Info(fmt.Printf(constants.ReadProcessDone, domainNameTask))

	result, err := repo.ReadTasksRepo()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("task not found"),
		})

		log.Error("task not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"tasks": result,
	})
	log.Info(fmt.Printf(constants.ReadProcessSuccess, domainNameTask))

	return
}

func UpdateTask(c *gin.Context) {
	log := initLog()
	log.Info(fmt.Printf(constants.UpdateProcess, domainNameTask))
	defer log.Info(fmt.Printf(constants.UpdateProcessDone, domainNameTask))

	var task modelDb.Task
	id := c.Param("id")
	err := c.ShouldBind(&task)
	task.TaskId = id

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		log.Error(err.Error())
		return
	}

	result, err := repo.UpdateTaskRepo(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		log.Error("task not updated")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"taks": result,
	})
	log.Info(fmt.Printf(constants.UpdateProcessSuccess, domainNameTask))

	return
}

func DeleteTask(c *gin.Context) {
	log := initLog()
	log.Info(fmt.Printf(constants.DeleteProcess, domainNameTask))
	defer log.Info(fmt.Printf(constants.DeleteProcessDone, domainNameTask))

	id := c.Param("id")
	_, err := repo.ReadTaskRepo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "task not found",
		})

		log.Error("task not found")
		return
	}

	err = repo.DeleteTaskRepo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		log.Error(fmt.Printf(constants.ErrDataNotFound, domainNameTask))

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "task deleted successfully",
	})
	log.Info(fmt.Printf(constants.DeleteProcessSuccess, domainNameTask))

	return
}
