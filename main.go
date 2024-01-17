package main

import (
	"PowerPlantManagementApplication/config"
	"PowerPlantManagementApplication/constants"
	"PowerPlantManagementApplication/controllers"
	"PowerPlantManagementApplication/routes"
	"PowerPlantManagementApplication/services"
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoClient *mongo.Client
	server      *gin.Engine
)

func initCompany(mongoClient *mongo.Client) {
	collection := config.GetCollection(mongoClient, constants.DatabaseName, "Company")
	companyService := services.Initservices(collection)
	companyController := controllers.Initcontroller(companyService)
	routes.Routes(server, companyController)
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

	initCompany(mongoClient)


	server.Run(constants.Port)
}
