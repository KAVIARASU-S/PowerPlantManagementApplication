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
	mongoClient *mongo.Client
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

func initTasks (mongoClient *mongo.Client){
	taskcollection := config.GetCollection(mongoClient, constants.DatabaseName, "Tasks")
	taskservice := services.InitTasks(taskcollection)
	taskController := controllers.InitTaskController(taskservice)
	routes.TaskRoutes(server,taskController)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := config.ConnectDatabase(ctx)
	defer mongoClient.Disconnect(ctx)

	if err != nil {
		log.Printf("Not Connected to Database! Resolve the issue!!!")
	}

	server = gin.Default()
	server.Use(corsMiddleware())

	initCompany(mongoClient)
	initUser(mongoClient)
	initTasks(mongoClient)

	server.Run(constants.Port)
}
