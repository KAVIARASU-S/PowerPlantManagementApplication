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

type SensorService struct {
	SensorCollection *mongo.Collection
}

func InitSensors(sensorCollection *mongo.Collection) interfaces.ISensor {
	return &SensorService{
		SensorCollection: sensorCollection,
	}
}

func (sensorData *SensorService) DisplaySensors(searchFilter models.SearchFilter) (allSensors *[]models.Sensors,err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"CompanyName":searchFilter.CompanyName}
	result, err := sensorData.SensorCollection.Find(ctx, filter)
	if err != nil {
		log.Println("Error Finding Items in mongoDB ", err)
		return nil, err
	}

	log.Println("Successfully found Items in mongoDb")

	var sensors []models.Sensors

	err = result.All(ctx, &sensors)

	if err != nil {
		log.Println("Error parsing the sensors to slice",err)
	}

	return &sensors, nil
}

func (sensorData *SensorService) InsertSensor(sensor *models.Sensors)(err error){
	log.Println("The sensor to be inserted is ",sensor)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result,err := sensorData.SensorCollection.InsertOne(ctx,sensor)

	if err != nil {
		log.Println("Error inserting sensor in mongoDB",err)
		return err
	}

	log.Println("Sucessfully added the sensor",result)

	return nil
}

func (sensorData *SensorService) UpdateSensor(sensor *models.Sensors)(err error){
	log.Println("The sensor to be updated is ",sensor)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"SensorName":sensor.SensorName}
	update := bson.D{{"$set",bson.D{
		{"LowerLimit",sensor.LowerLimit},
		{"UpperLimit",sensor.UpperLimit},
	}}}
	result,err := sensorData.SensorCollection.UpdateOne(ctx,filter,update)

	if err != nil {
		log.Println("Error when updating sensor", err)
		return err
	}

	log.Println("Sucessfully Updated sensor", result)

	return nil
}

func (sensorData *SensorService) DeleteSensor(sensor *models.Sensors)(err error){
	log.Println("The sensor to be deleted is ",sensor)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result,err := sensorData.SensorCollection.DeleteOne(ctx,sensor)

	if err != nil {
		log.Println("Error when deleting item", err)
		return err
	}

	log.Println("Successfully Deleted the item", result)
	return nil
}