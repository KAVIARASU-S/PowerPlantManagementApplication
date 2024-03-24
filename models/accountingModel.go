package models

import "time"

type Transaction struct {
	Amount int `json:"Amount" bson:"Amount"`
	Date   time.Time `json:"Date" bson:"Date"`
	TransactionType string `json:"TransactionType" bson:"TransactionType"`
	Description string `json:"Description" bson:"Description"`
}
