package interfaces

import "PowerPlantManagementApplication/models"

type Icompany interface {
	DisplayCompany () (*[]models.Company,error)
	InsertCompany (company *models.Company) (err error)
}