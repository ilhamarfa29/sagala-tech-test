package task

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"slices"

	modelDb "sagala-tech-test/database/model"
	repo "sagala-tech-test/internal/app/task/repository"
	constants "sagala-tech-test/internal/constant"

	"github.com/gin-gonic/gin"
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
	log.Info(fmt.Sprintf(constants.CreateProcess, domainNameTask))
	defer log.Info(fmt.Sprintf(constants.CreateProcessDone, domainNameTask))

	var task *modelDb.Task
	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		log.Error(err.Error())
		return
	}

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

	log.Info(fmt.Sprintf(constants.CreateProcessSuccess, domainNameTask))

	return
}

func ReadTask(c *gin.Context) {
	log := initLog()
	log.Info(fmt.Sprintf(constants.ReadByIdProcess, domainNameTask))
	defer log.Info(fmt.Sprintf(constants.ReadByIdProcessDone, domainNameTask))

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
	log.Info(fmt.Sprintf(constants.ReadByIdProcessSuccess, domainNameTask))

	return
}

func ReadTasks(c *gin.Context) {
	log := initLog()
	log.Info(fmt.Sprintf(constants.ReadProcess, domainNameTask))
	defer log.Info(fmt.Sprintf(constants.ReadProcessDone, domainNameTask))

	var filter *modelDb.Filter
	err := c.ShouldBind(&filter)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("error parsing payload"),
		})

		log.Error("error parsing payload")
		return
	}

	status := c.Param("status")
	statuses := []string{"waiting_list", "in_progress", "done", ""}

	statusAvailable := slices.Contains(statuses, status)
	if !statusAvailable {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "status " + status + " not found",
		})

		log.Error("status " + status + " not found")
		return
	}

	result, err := repo.ReadTasksRepo(filter)
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
	log.Info(fmt.Sprintf(constants.ReadProcessSuccess, domainNameTask))

	return
}

func UpdateTask(c *gin.Context) {
	log := initLog()
	log.Info(fmt.Sprintf(constants.UpdateProcess, domainNameTask))
	defer log.Info(fmt.Sprintf(constants.UpdateProcessDone, domainNameTask))

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

	resultTask, err := repo.ReadTaskRepo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "task not found",
		})

		log.Error("task not found")
		return
	}

	if resultTask.Status != task.Status {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "update task cannot use for update status",
		})

		log.Error("update task cannot use for update status")
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
		"task": result,
	})
	log.Info(fmt.Sprintf(constants.UpdateProcessSuccess, domainNameTask))

	return
}

func UpdateStatusTask(c *gin.Context) {
	log := initLog()
	log.Info(fmt.Sprintf(constants.UpdateStatusProcess, domainNameTask))
	defer log.Info(fmt.Sprintf(constants.UpdateStatusProcessDone, domainNameTask))

	id := c.Param("id")
	status := c.Param("status")
	statuses := []string{"waiting_list", "in_progress", "done"}

	statusAvailable := slices.Contains(statuses, status)
	if !statusAvailable {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "status " + status + " not found",
		})

		log.Error("status " + status + " not found")
		return
	}

	resultTask, err := repo.ReadTaskRepo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "task not found",
		})

		log.Error("task not found")
		return
	}

	resultTask.Status = status

	result, err := repo.UpdateTaskRepo(resultTask)
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
	log.Info(fmt.Sprintf(constants.UpdateStatusProcessDone, domainNameTask))

	return
}

func SoftDeleteTask(c *gin.Context) {
	log := initLog()
	log.Info(fmt.Sprintf(constants.UpdateProcess, domainNameTask))
	defer log.Info(fmt.Sprintf(constants.UpdateProcessDone, domainNameTask))

	id := c.Param("id")
	resultTask, err := repo.ReadTaskRepo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		log.Error(err.Error())
		return
	}

	resultTask.IsDeleted = true

	result, err := repo.UpdateTaskRepo(resultTask)
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
	log.Info(fmt.Sprintf(constants.UpdateProcessSuccess, domainNameTask))

	return
}

func DeleteTask(c *gin.Context) {
	log := initLog()
	log.Info(fmt.Sprintf(constants.DeleteProcess, domainNameTask))
	defer log.Info(fmt.Sprintf(constants.DeleteProcessDone, domainNameTask))

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
		log.Error(fmt.Sprintf(constants.ErrDataNotFound, domainNameTask))

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "task deleted successfully",
	})
	log.Info(fmt.Sprintf(constants.DeleteProcessSuccess, domainNameTask))

	return
}
