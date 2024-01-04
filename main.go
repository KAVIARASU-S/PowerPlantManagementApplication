package main

import (
	"PowerPlantManagementApplication/models"
	"fmt"
)

func main() {
	var Company models.Company
	Company.CompanyName = "ford"
	fmt.Println(Company)
}
