package models

type Sensors struct{
	SensorName string `json:"SensorName" bson:"SensorName"`
	LowerLimit int `json:"LowerLimit" bson:"LowerLimit"`
	UpperLimit int `json:"UpperLimit" bson:"UpperLimit"`
}