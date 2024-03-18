package interfaces

import "PowerPlantManagementApplication/models"

type ISensor interface {
	DisplaySensors() (allSensors *[]models.Sensors, err error)
	InsertSensor(sensor *models.Sensors)(err error)
	UpdateSensor(sensor *models.Sensors)(err error)
	DeleteSensor(sensor *models.Sensors)(err error)
}
