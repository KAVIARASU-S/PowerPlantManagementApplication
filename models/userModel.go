package models


type User struct {
	UserName string `json:"UserName" bson:"UserName"`
	Password string `json:"Password" bson:"Password"`
	Company string `json:"Company" bson:"Company"`
	Role string `json:"Role" bson:"Role"`
	PowerplantType string `json:"PowerplantType" bson:"PowerplantType"`
	Totp string `json:"totp" bson:"totp"`
}

type IPAddress struct {
	IPaddress string `json:"IPaddress" bson:"IPaddress"`
}

type Login struct{
	UserName string `json:"UserName" bson:"UserName"`
	Password string `json:"Password" bson:"Password"`
	Totp string `json:"totp" bson:"totp"`
}

type ShowUser struct{
	UserName string `json:"UserName" bson:"UserName"`
	Company string `json:"Company" bson:"Company"`
	Role string `json:"Role" bson:"Role"`
	PowerplantType string `json:"PowerplantType" bson:"PowerplantType"`
}