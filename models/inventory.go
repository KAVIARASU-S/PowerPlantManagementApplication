package models

import "time"

type Inventory struct{
	CompanyName string `json:"CompanyName" bson:"CompanyName"`
	ItemName string `json:"ItemName" bson:"ItemName"`
	Brand string `json:"Brand" bson:"Brand"`
	Price int `json:"Price" bson:"Price"`
	Quantity int `json:"Quantity" bson:"Quantity"`
	ReplacementDate time.Time `json:"ReplacementDate" bson:"ReplacementDate"`
}