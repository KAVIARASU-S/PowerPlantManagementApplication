package interfaces

import "PowerPlantManagementApplication/models"

type Iuser interface{
	CreateUser (user *models.User) (qr string,err error)
	ValidateUser (Login *models.Login) (err error)
	ValidateIP (ip *models.IPAddress) (err error)
	InsertIP (ip *models.IPAddress) (err error)
	ValidateTotp (user *models.Login) (company string,role string,plantType string,err error)
	DisplayUser () (allusers *[]models.ShowUser,err error)
	DisplayIP () (allusers *[]models.IPAddress,err error)
}