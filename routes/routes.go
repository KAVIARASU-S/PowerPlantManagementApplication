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

func TaskRoutes (router *gin.Engine, controller controllers.TaskController){
	router.GET("/displayTasks",controller.DisplayTasks)
	router.POST("/insertTask",controller.InsertTask)
	router.POST("/updateTask",controller.UpdateTask)
	router.POST("/deleteTask",controller.DeleteTask)
}

func InventoryRoutes (router *gin.Engine, controller controllers.InventoryController){
	router.GET("/displayItems",controller.DisplayItems)
	router.POST("/addItem",controller.AddItem)
	router.POST("/updateItem",controller.UpdateItem)
	router.POST("/deleteItem",controller.DeleteItem)
	router.GET("/displayPurchase",controller.DisplayPurchase)
	router.POST("/addPurchase",controller.AddPurchase)
	router.POST("/deletePurchase",controller.DeletePurchase)
}

func SensorRoutes (router *gin.Engine,controller controllers.SensorController){
	router.GET("/displaySensors",controller.DisplaySensors)
	router.POST("/insertSensor",controller.InsertSensor)
	router.POST("/updateSensor",controller.UpdateSensor)
	router.POST("/deleteSensor",controller.DeleteSensor)
}
