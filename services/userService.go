package services

import (
	"PowerPlantManagementApplication/interfaces"
	"PowerPlantManagementApplication/models"
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceModel struct{
	UserCollection *mongo.Collection
	IPCollection *mongo.Collection
}

func InitUserService(usercollection,ipcollection *mongo.Collection) interfaces.Iuser {
	return &UserServiceModel{
		UserCollection: usercollection,
		IPCollection: ipcollection,
	}
}

func (userData *UserServiceModel)CreateUser (user *models.User) (err error){
	log.Println("Data entered by user",user)


	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	check := userData.UserCollection.FindOne(ctx,bson.M{"UserName": user.UserName})

	if check.Err() == nil {
		log.Println("The user you are trying to create already exists")
		error := errors.New("User already exists")
		return error
	}

	result, err := userData.UserCollection.InsertOne(ctx, user)

	if err != nil {
		log.Println("Error inserting to MongoDB: ", err)
		return err
	} else {
		log.Println("User Inserted to MongoDb successfully")
	}

	log.Println("Successfully inserted",result)

	return nil
}

func (userData *UserServiceModel)ValidateUser (user *models.User) (err error){
	log.Println("Data entered by user to be validated",user)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result:= userData.UserCollection.FindOne(ctx,user)

	if result.Err() != nil {
		log.Println("Error validating user: ", err)
		error := errors.New("Wrong User name and password")
		return error
	} else {
		log.Println("User found successfully")
	}

	var validatedUser models.User

	if err := result.Decode(&validatedUser); err != nil {
		log.Println("Error decoding",err)
	}

	log.Println("Successfully Validated")
	log.Println("The user is ",validatedUser)
	return nil
}

func(userData *UserServiceModel)InsertIP(ip *models.IPAddress) (err error){
	log.Println("IP address to be added: ",ip)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := userData.IPCollection.InsertOne(ctx, ip)

	if err != nil {
		log.Println("Error inserting to MongoDB: ", err)
		return err
	} else {
		log.Println("IP Inserted to MongoDb successfully")
	}

	log.Println("Successfully inserted",result)

	return nil
}

func (userData *UserServiceModel) ValidateIP(ip *models.IPAddress) (err error){
	log.Println("IP address of the user: ",ip)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result:= userData.IPCollection.FindOne(ctx,ip)

	if result.Err() != nil {
		log.Println("Error validating IP: ", err)
		error := errors.New("Wrong IP Address")
		return error
	} else {
		log.Println("IP Address validated successfully")
	}

	return nil
}