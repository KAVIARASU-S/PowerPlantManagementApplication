package services

import (
	"PowerPlantManagementApplication/interfaces"

	"go.mongodb.org/mongo-driver/mongo"
)

type InventoryServiceModel struct {
	InventoryCollection *mongo.Collection
}

func InitInventory (collection *mongo.Collection)(interfaces.IInventory){
	return &InventoryServiceModel{
		InventoryCollection: collection,
	}
}

func 