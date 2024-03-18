package services

import (
	"PowerPlantManagementApplication/interfaces"
	"PowerPlantManagementApplication/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CompanyServiceModel struct {
	Collection *mongo.Collection
}

func Initservices(collection *mongo.Collection) interfaces.Icompany {
	return &CompanyServiceModel{Collection: collection}
}

func (companyData *CompanyServiceModel) DisplayCompany () (*[]models.Company,error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := companyData.Collection.Find(ctx, bson.D{})

	if err != nil {
		log.Println("Error finding data in MongoDB: ", err)
		return nil,err
	} else {
		log.Println("Found data in MongoDb successfully")
	}

	var results []models.Company

	err = result.All(ctx, &results)

	if err != nil {
		log.Println("Error parsing the companies to slice",err)
	}

	for _, i := range results {
		log.Println(i)
	}

	return &results,nil
}


func (companyData *CompanyServiceModel) InsertCompany(company *models.Company) (err error) {
	log.Println("Data entered by user",company)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := companyData.Collection.InsertOne(ctx, company)

	if err != nil {
		log.Println("Error inserting to MongoDB: ", err)
		return err
	} else {
		log.Println("Inserted to MongoDb successfully")
	}

	log.Println("Successfully inserted",result)

	return nil
}

