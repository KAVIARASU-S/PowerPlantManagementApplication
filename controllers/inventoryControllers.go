package controllers

import (
	"PowerPlantManagementApplication/interfaces"

	"github.com/gin-gonic/gin"
)


type InventoryController struct{
	InventoryService interfaces.IInventory
}

func InitInventoryController (inventoryService interfaces.IInventory)(InventoryController){
	return InventoryController{InventoryService: inventoryService}
}

func(controller *InventoryController)DisplayItems(c *gin.Engine){
	
}