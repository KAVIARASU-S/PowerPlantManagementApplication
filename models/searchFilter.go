package models

type SearchFilter struct{
	CompanyName string `json:"CompanyName" bson:"CompanyName"`
}