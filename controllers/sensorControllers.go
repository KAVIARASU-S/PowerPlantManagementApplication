package controllers

import (
	"PowerPlantManagementApplication/interfaces"
	"PowerPlantManagementApplication/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SensorController struct {
	SensorService interfaces.ISensor
}

func InitSensorController(sensorservice interfaces.ISensor) SensorController {
	return SensorController{
		SensorService: sensorservice,
	}
}

func (controller *SensorController) DisplaySensors(c *gin.Context) {
	allSensors, err := controller.SensorService.DisplaySensors()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, allSensors)
}

func (controller *SensorController) InsertSensor(c *gin.Context) {
	var sensor *models.Sensors

	if err := c.ShouldBindJSON(&sensor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Json format",
		})
		return
	}

	if err := controller.SensorService.InsertSensor(sensor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": "Sensor inserted successfully",
	})

}

func (controller *SensorController) UpdateSensor(c *gin.Context) {
	var sensor *models.Sensors

	if err := c.ShouldBindJSON(&sensor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Json format",
		})
		return
	}

	if err := controller.SensorService.UpdateSensor(sensor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": "Sensor updated successfully",
	})
}

func (controller *SensorController) DeleteSensor(c *gin.Context) {
	var sensor *models.Sensors

	if err := c.ShouldBindJSON(&sensor); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Json format",
		})
		return
	}

	if err := controller.SensorService.DeleteSensor(sensor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": "Sensor deleted successfully",
	})
}
