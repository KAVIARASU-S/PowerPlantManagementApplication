package models

type User struct {
	UserName string `json:"UserName" bson:"UserName"`
	Password string `json:"Password" bson:"Password"`
}

type IPAddress struct {
	IPaddress string `json:"IPaddress" bson:"IPaddress"`
}