package interfaces

import "PowerPlantManagementApplication/models"

type IInventory interface {
	DisplayItems(searchFilter *models.SearchFilter) (allitems *[]models.Inventory, err error)
	AddItem(item *models.Inventory) (err error)
	UpdateItem(item *models.Inventory)(err error)
	DeleteItem(item *models.Inventory) (err error)
	DisplayPurchase(searchFilter *models.SearchFilter) (allpurchase *[]models.Inventory,err error)
	AddPurchase(item *models.Inventory) (err error)
	DeletePurchase(item *models.Inventory) (err error)
}
