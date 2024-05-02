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
	var searchFilter models.SearchFilter

	if err := c.ShouldBindJSON(&searchFilter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Json format",
		})
		return
	}

	allItems, err := controller.InventoryService.DisplayItems(&searchFilter)

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
	var searchFilter models.SearchFilter

	if err := c.ShouldBindJSON(&searchFilter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Json format",
		})
		return
	}

	allPurchase, err := controller.InventoryService.DisplayPurchase(&searchFilter)

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
