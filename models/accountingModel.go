package models

import "time"

type Transaction struct {
	CompanyName string `json:"CompanyName" bson:"CompanyName"`
	Amount int `json:"Amount" bson:"Amount"`
	Date   time.Time `json:"Date" bson:"Date"`
	TransactionType string `json:"TransactionType" bson:"TransactionType"`
	Description string `json:"Description" bson:"Description"`
}

type Accounting struct {
	CompanyName string `json:"CompanyName" bson:"CompanyName"`
	Incomes []Transaction `json:"Incomes" bson:"Incomes"`
	Expenses []Transaction `json:"Expenses" bson:"Expenses"`
	TotalIncome int `json:"TotalIncome" bson:"TotalIncome"`
	TotalExpense int `json:"TotalExpense" bson:"TotalExpense"`
	Profit int `json:"Profit" bson:"Profit"`
}