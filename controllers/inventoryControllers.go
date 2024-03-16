package controllers

import (
	"PowerPlantManagementApplication/interfaces"
	"PowerPlantManagementApplication/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InventoryController struct {
	InventoryService interfaces.IInventory
}

func InitInventoryController(inventoryService interfaces.IInventory) InventoryController {
	return InventoryController{InventoryService: inventoryService}
}

func (controller *InventoryController) DisplayItems(c *gin.Context) {
	allItems, err := controller.InventoryService.DisplayItems()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, allItems)
}

func (controller *InventoryController) AddItem(c *gin.Context) {
	var item *models.Inventory

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Json format",
		})
		return
	}

	if err := controller.InventoryService.AddItem(item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": "Item inserted successfully",
	})
}

func (controller *InventoryController) UpdateItem(c *gin.Context) {
	var item *models.Inventory

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Json format",
		})
		return
	}

	if err := controller.InventoryService.UpdateItem(item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": "Item updated successfully",
	})
}

func (controller *InventoryController) DeleteItem(c *gin.Context) {
	var item *models.Inventory

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Json format",
		})
		return
	}

	if err := controller.InventoryService.DeleteItem(item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": "Item deleted successfully",
	})
}

func (controller *InventoryController) DisplayPurchase(c *gin.Context) {
	allPurchase, err := controller.InventoryService.DisplayPurchase()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, allPurchase)
}

func (controller *InventoryController) AddPurchase(c *gin.Context) {
	var item *models.Inventory

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Json format",
		})
		return
	}

	if err := controller.InventoryService.AddPurchase(item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": "Purchase Item added successfully",
	})

}

func (controller *InventoryController) DeletePurchase(c *gin.Context) {
	var item *models.Inventory

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Json format",
		})
		return
	}

	if err := controller.InventoryService.DeletePurchase(item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": "Purchase Item deleted successfully",
	})
}
