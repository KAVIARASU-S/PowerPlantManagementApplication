package controllers

import (
	"PowerPlantManagementApplication/interfaces"
	"PowerPlantManagementApplication/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskService interfaces.ITasks
}

func InitTaskController (taskservice interfaces.ITasks) (TaskController){
	return TaskController{TaskService: taskservice}
}

func (controller *TaskController)DisplayTasks(c *gin.Context){
	allTasks,err := controller.TaskService.DisplayTasks()

	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"Error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,allTasks)
}

func (controller *TaskController) InsertTask(c *gin.Context){
	var task *models.Tasks

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"Error":"Invalid Json format",
		})
		return
	}

	if err := controller.TaskService.InsertTask(task); err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"Error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"Status":"Task inserted successfully",
	})

}

func (controller *TaskController) UpdateTask (c *gin.Context){
	var task *models.Tasks

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"Error":"Invalid Json format",
		})
		return
	}

	if err := controller.TaskService.UpdateTask(task); err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"Error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"Status":"Task updated successfully",
	})

}

func (controller *TaskController) DeleteTask (c *gin.Context){
	var task *models.Tasks

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"Error":"Invalid Json format",
		})
		return
	}

	if err := controller.TaskService.DeleteTask(task); err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"Error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"Status":"Task Deleted successfully",
	})
}