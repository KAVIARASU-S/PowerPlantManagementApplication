package interfaces

import "PowerPlantManagementApplication/models"

type IAccounting interface {
	DisplayTransactions(searchFilter *models.SearchFilter)(allTransactions *[]models.Transaction,err error)
	InsertTransaction(transaction *models.Transaction)(err error)
	DisplayAccounting(searchFilter *models.SearchFilter)(allAccounts *models.Accounting,err error)
}
