package models

import "time"

type Inventory struct{
	ItemName string `json:"ItemName" bson:"Itemname"`
	Brand string `json:"Brand" bson:"Brand"`
	Price string `json:"Price" bson:"Price"`
	Quantity string `json:"Quantity" bson:"Quantity"`
	ReplacementDate time.Time `json:"ReplacementDate" bson:"ReplacementDate"`
}