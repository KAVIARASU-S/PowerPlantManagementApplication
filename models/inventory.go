package models

import "time"

type Inventory struct{
	ItemName string
	Brand string
	Price string
	Quantity string
	ReplacementDate time.Time
}