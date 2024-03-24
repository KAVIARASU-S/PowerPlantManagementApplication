package interfaces

import "PowerPlantManagementApplication/models"

type IAccounting interface {
	DisplayTransactions()(allTransactions *[]models.Transaction,err error)
	InsertTransaction(transaction *models.Transaction)(err error)
}
