package main

import (
	"PowerPlantManagementApplication/config"
	"PowerPlantManagementApplication/constants"
	"PowerPlantManagementApplication/controllers"
	"PowerPlantManagementApplication/routes"
	"PowerPlantManagementApplication/services"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server      *gin.Engine
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func initCompany(mongoClient *mongo.Client) {
	collection := config.GetCollection(mongoClient, constants.DatabaseName, "Company")
	companyService := services.Initservices(collection)
	companyController := controllers.Initcontroller(companyService)
	routes.Routes(server, companyController)
}

func initUser(mongoClient *mongo.Client){
	usercollection := config.GetCollection(mongoClient, constants.DatabaseName, "User")
	ipcollection := config.GetCollection(mongoClient,constants.DatabaseName,"IPAddress")
	userService := services.InitUserService(usercollection,ipcollection)
	
	userController := controllers.InitUserController(userService)
	routes.UserRoutes(server, userController)
}

func initTasks(mongoClient *mongo.Client){
	taskcollection := config.GetCollection(mongoClient, constants.DatabaseName, "Tasks")
	taskservice := services.InitTasks(taskcollection)
	taskController := controllers.InitTaskController(taskservice)
	routes.TaskRoutes(server,taskController)
}

func initInventory(mongoClient *mongo.Client){
	inventorycollection := config.GetCollection(mongoClient, constants.DatabaseName, "Inventory")
	purchasecollection := config.GetCollection(mongoClient, constants.DatabaseName, "Purchase")
	inventoryservice := services.InitInventory(inventorycollection,purchasecollection)
	inventoryController := controllers.InitInventoryController(inventoryservice)
	routes.InventoryRoutes(server,inventoryController)
}

func initSensors(mongoClient *mongo.Client){
	sensorcollection := config.GetCollection(mongoClient, constants.DatabaseName, "Sensors")
	sensorservice := services.InitSensors(sensorcollection)
	sensorController := controllers.InitSensorController(sensorservice)
	routes.SensorRoutes(server,sensorController)
}

func initRoutes(mongoClient *mongo.Client){
	initCompany(mongoClient)
	initUser(mongoClient)
	initTasks(mongoClient)
	initInventory(mongoClient)
	initSensors(mongoClient)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := config.ConnectDatabase(ctx)
	defer mongoClient.Disconnect(ctx)

	if err != nil {
		log.Printf("Not Connected to Database! Resolve the issue!!!")
		log.Println(err)
	}

	server = gin.Default()
	server.Use(corsMiddleware())

	initRoutes(mongoClient)

	server.Run(constants.Port)
}
