package models


type Tasks struct {
	Task string `json:"Task" bson:"Task"`
	Deadline string `json:"Deadline" bson:"Deadline"`
	Employee string `json:"Employee" bson:"Employee"`
	Status bool `json:"Status" bson:"Status"`
}