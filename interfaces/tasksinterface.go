package interfaces

import "PowerPlantManagementApplication/models"

type ITasks interface{
	DisplayTasks (searchFilter *models.SearchFilter)(alltasks *[]models.Tasks,err error)
	InsertTask (task *models.Tasks)(err error)
	UpdateTask (task *models.Tasks)(err error)
	DeleteTask (task *models.Tasks) (err error)
}