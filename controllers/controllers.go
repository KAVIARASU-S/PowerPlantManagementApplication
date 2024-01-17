package controllers

import (
	"PowerPlantManagementApplication/interfaces"
	"PowerPlantManagementApplication/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompanyController struct {
	CompanyService interfaces.Icompany
}

func Initcontroller (companyService interfaces.Icompany) (CompanyController){
	return CompanyController{
		CompanyService: companyService,
	}
}

func (controller *CompanyController) DisplayCompany (c *gin.Context){
	allCompany,err := controller.CompanyService.DisplayCompany()

	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"Error in displaying Company Names":err,
		})
		return
	}
	
	c.JSON(http.StatusOK,allCompany)
}

func (controller *CompanyController) InsertCompany (c *gin.Context){
	var company *models.Company

	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"Error":"Invalid JSON format",
		})
		return
	}

	if err := controller.CompanyService.InsertCompany(company); err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"Error":err,
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"Status":"Success",
		"Inserted":company,
	})

}