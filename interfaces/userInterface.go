package interfaces

import "PowerPlantManagementApplication/models"

type Iuser interface{
	CreateUser (user *models.User) (err error)
	ValidateUser (user *models.User) (err error)
	ValidateIP (ip *models.IPAddress) (err error)
	InsertIP (ip *models.IPAddress) (err error)
}