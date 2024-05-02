package services

import (
	"PowerPlantManagementApplication/interfaces"
	"PowerPlantManagementApplication/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskServiceModel struct {
	TaskCollection *mongo.Collection
}

func InitTasks(collection *mongo.Collection)(interfaces.ITasks){
	return &TaskServiceModel{
		TaskCollection: collection,
	}
}

func(taskdata *TaskServiceModel) DisplayTasks (searchFilter *models.SearchFilter)(alltasks *[]models.Tasks,err error){
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"CompanyName":searchFilter.CompanyName}
	opts := options.Find().SetSort(map[string]int{"Deadline": 1})

	result,err:=taskdata.TaskCollection.Find(ctx,filter,opts)

	if err != nil {
		log.Println("Error finding data in MongoDB: ", err)
		return nil,err
	} else {
		log.Println("Found data in MongoDb successfully")
	}
	
	var tasks []models.Tasks

	result.All(ctx,&tasks)
	if err != nil {
		log.Println("Error parsing the tasks to slice",err)
	}

	log.Println("All tasks returned successfully")

	return &tasks,nil

}

func (taskdata *TaskServiceModel)InsertTask (task *models.Tasks)(err error){
	log.Println("The task to be inserted is ",task)

	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	result,err := taskdata.TaskCollection.InsertOne(ctx,task)

	if err != nil {
		log.Println("Error inserting Tasks in mongoDB",err)
		return err
	}

	log.Println("Successfully inserted Task",result)

	return nil 
}

func (taskdata *TaskServiceModel) UpdateTask (task *models.Tasks)(err error){
	log.Println("The task to be updated is ",task)

	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	filter := bson.D{
		{"$and",bson.A{
		bson.D{{"Task",task.Task}},
		bson.D{{"Employee",task.Employee}},
	}},
}

    update := bson.D{
		{"$set",bson.D{
			{"Status",task.Status},
		}},
	}

	result,err := taskdata.TaskCollection.UpdateOne(ctx,filter,update)

	if err != nil {
		log.Println("Error when updating task")
		return err
	}

	log.Println("Successfully updated Task",result)

	return nil 
}

func (taskdata *TaskServiceModel) DeleteTask (task *models.Tasks) (err error){
	log.Println("The task to be Deleted is ", task)
	
	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	filter := bson.D{
		{"$and",bson.A{
			bson.D{{"Task",task.Task}},
			bson.D{{"Employee",task.Employee}},
		}},
	}

	result,err := taskdata.TaskCollection.DeleteOne(ctx,filter)

	if err != nil {
		log.Println("Error Deleting the task",err)
		return err
	}

	log.Println("Successfully deleted the task ",result)

	return nil
}