package controllers

import (
	"PowerPlantManagementApplication/interfaces"
	"PowerPlantManagementApplication/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountingController struct {
	AccountingService interfaces.IAccounting
}

func InitAccountingController(acountingService interfaces.IAccounting) AccountingController {
	return AccountingController{
		AccountingService: acountingService,
	}
}

func (controller *AccountingController) DisplayTransactions(c *gin.Context) {
	var searchFilter models.SearchFilter

	if err := c.ShouldBindJSON(&searchFilter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Json format",
		})
		return
	}

	allTransactions, err := controller.AccountingService.DisplayTransactions(&searchFilter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, allTransactions)
}

func (controller *AccountingController) InsertTransaction(c *gin.Context) {
	var transaction *models.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Json format",
		})
		return
	}

	if err := controller.AccountingService.InsertTransaction(transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": "Transaction inserted successfully",
	})
}

func (controller *AccountingController) DisplayAccounting(c *gin.Context){
	var searchFilter models.SearchFilter

	if err := c.ShouldBindJSON(&searchFilter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid Json format",
		})
		return
	}

	allAccounts,err := controller.AccountingService.DisplayAccounting(&searchFilter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, allAccounts)
}
