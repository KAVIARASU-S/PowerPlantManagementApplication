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

type InventoryServiceModel struct {
	InventoryCollection *mongo.Collection
	PurchaseCollection *mongo.Collection
}

func InitInventory(inventoryCollection *mongo.Collection,purchaseCollection *mongo.Collection) interfaces.IInventory {
	return &InventoryServiceModel{
		InventoryCollection: inventoryCollection,
		PurchaseCollection: purchaseCollection,
	}
}

func (inventoryData *InventoryServiceModel) DisplayItems(searchFilter *models.SearchFilter) (allitems *[]models.Inventory, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"CompanyName":searchFilter.CompanyName}
	opts := options.Find().SetSort(map[string]int{"ReplacementDate": 1})

	result, err := inventoryData.InventoryCollection.Find(ctx, filter, opts)
	if err != nil {
		log.Println("Error Finding Items in mongoDB ", err)
		return nil, err
	}

	log.Println("Successfully found Items in mongoDb")

	var items []models.Inventory

	err = result.All(ctx, &items)

	if err != nil {
		log.Println("Error parsing the items to slice",err)
	}

	return &items, nil
}

func (inventoryData *InventoryServiceModel) AddItem(item *models.Inventory) (err error) {
	log.Println("Item to be Added ", item)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := inventoryData.InventoryCollection.InsertOne(ctx, item)

	if err != nil {
		log.Println("Error inserting Item in mongoDB",err)
		return err
	}

	log.Println("Successfully Inserted Item", result)

	return nil
}

func (inventoryData *InventoryServiceModel) UpdateItem(item *models.Inventory) (err error) {
	log.Println("Item to be updated ", item)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"ItemName": item.ItemName,
		"Brand":    item.Brand,
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "Quantity", Value: item.Quantity},
		}},
	}

	result, err := inventoryData.InventoryCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		log.Println("Error when updating item", err)
		return err
	}

	log.Println("Sucessfully Updated Item", result)

	return nil
}

func (inventoryData *InventoryServiceModel) DeleteItem(item *models.Inventory) (err error) {
	log.Println("The item to be deleted is ", item)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"ItemName": item.ItemName,
		"Brand":    item.Brand,
	}
	result, err := inventoryData.InventoryCollection.DeleteOne(ctx, filter)

	if err != nil {
		log.Println("Error when deleting item", err)
		return err
	}

	log.Println("Successfully Deleted the item", result)
	return nil
}

func (inventoryData *InventoryServiceModel) DisplayPurchase(searchFilter *models.SearchFilter) (allpurchase *[]models.Inventory,err error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"CompanyName":searchFilter.CompanyName}
	result,err := inventoryData.PurchaseCollection.Find(ctx,filter)
	if err != nil {
		log.Println("Error Finding Purchase Items in mongoDB ", err)
		return nil, err
	}

	log.Println("Successfully found Purchase Items in mongoDb")

	var purchase []models.Inventory

	result.All(ctx, &purchase)
	return &purchase,nil
}


func (inventoryData *InventoryServiceModel) AddPurchase(item *models.Inventory) (err error){
	log.Println("The item to be added for purchase is ", item)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := inventoryData.PurchaseCollection.InsertOne(ctx, item)

	if err != nil {
		log.Println("Error inserting Purchase Item in mongoDB",err)
		return err
	}

	log.Println("Successfully Inserted Purchase Item", result)

	return nil 
}

func (inventoryData *InventoryServiceModel) DeletePurchase(item *models.Inventory) (err error){
	log.Println("The item to be deleted form purchase is ", item)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"ItemName": item.ItemName,
		"Brand":    item.Brand,
	}
	result, err := inventoryData.PurchaseCollection.DeleteOne(ctx,filter)

	if err != nil {
		log.Println("Error deleting Purchase Item in mongoDB")
		return err
	}

	log.Println("Successfully Deleted Purchase Item", result)

	return nil 
}