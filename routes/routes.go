package routes

import (
	"PowerPlantManagementApplication/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine,controller controllers.CompanyController) {
	router.GET("/sample",controller.DisplayCompany)
	router.POST("/sample",controller.InsertCompany)
}
