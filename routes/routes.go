package routes

import (
	"PowerPlantManagementApplication/controllers"

	"github.com/gin-gonic/gin"
)



func Routes(router *gin.Engine,controller controllers.CompanyController) {
	router.GET("/sample",controller.DisplayCompany)
	router.POST("/sample",controller.InsertCompany)
}

func UserRoutes (router *gin.Engine,controller controllers.UserController){
	router.POST("/createUser",controller.CreateUser)
	router.POST("/signIn",controller.ValidateUser)
	router.POST("/insertIP",controller.InsertIP)
	router.POST("/totp",controller.ValidateTotp)
	router.GET("/displayUser",controller.DisplayUser)
	router.GET("/displayIP",controller.DisplayIP)
}