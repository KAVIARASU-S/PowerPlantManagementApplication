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

type AccountingService struct {
	AccountingCollection *mongo.Collection
}

func InitAccounting (accountingCollection *mongo.Collection)interfaces.IAccounting{
	return &AccountingService{
		AccountingCollection: accountingCollection,
	}
}

func (accountingData *AccountingService) DisplayTransactions()(allTransactions *[]models.Transaction,err error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Find().SetSort(map[string]int{"Date": -1})

	result, err := accountingData.AccountingCollection.Find(ctx, bson.M{},opts)
	if err != nil {
		log.Println("Error Finding transactions in mongoDB ", err)
		return nil, err
	}

	log.Println("Successfully found Transactions in mongoDb")

	var transactions []models.Transaction

	err = result.All(ctx, &transactions)

	if err != nil {
		log.Println("Error parsing the Transactions to slice",err)
	}

	return &transactions,nil
}

func (accountingData *AccountingService) InsertTransaction(transaction *models.Transaction)(err error){
	log.Println("Transaction to be inserted ", transaction)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := accountingData.AccountingCollection.InsertOne(ctx, transaction)

	if err != nil {
		log.Println("Error inserting transaction in mongoDB",err)
		return err
	}

	log.Println("Successfully Inserted Transaction", result)
	return nil
}