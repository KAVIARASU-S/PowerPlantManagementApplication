package services

import (
	"PowerPlantManagementApplication/interfaces"
	"PowerPlantManagementApplication/models"
	"context"
	"encoding/base64"
	"errors"
	"log"
	"time"

	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceModel struct{
	UserCollection *mongo.Collection
	IPCollection *mongo.Collection
}

func InitUserService(usercollection,ipcollection *mongo.Collection) interfaces.Iuser {
	return &UserServiceModel{
		UserCollection: usercollection,
		IPCollection: ipcollection,
	}
}

func (userData *UserServiceModel)CreateUser (user *models.User) (qr string,err error){
	log.Println("Data entered by user",user)


	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	check := userData.UserCollection.FindOne(ctx,bson.M{"UserName": user.UserName})

	if check.Err() == nil {
		log.Println("The user you are trying to create already exists")
		error := errors.New("username already exists")
		return "",error
	}

	Totpcode, qr, err := generateTotp(user) 
	if err != nil {
		log.Println("Error generating totp and qr : ", err)
		return "",err
	}

	user.Totp = Totpcode

	result, err := userData.UserCollection.InsertOne(ctx, user)

	if err != nil {
		log.Println("Error inserting to MongoDB: ", err)
		return "",err
	} else {
		log.Println("User Inserted to MongoDb successfully")
	}

	log.Println("Successfully inserted",result)

	return qr,nil
}

func generateTotp(user *models.User) (Totpcode string,qr string,err error){
	key ,err := totp.Generate(totp.GenerateOpts{
		Issuer: "Power Plant",
		AccountName: user.UserName,
	})

	if err != nil {
		log.Println("Error in generating totp: ",err)
		return "","",err
	}

	otpAuthURL := key.URL()

	totpSecret := key.Secret()

	log.Println("TOTP secret for ",user.UserName, " : ",totpSecret)

	qrCode, err := qrcode.Encode(otpAuthURL,qrcode.Medium,256)

	if err != nil {
		log.Println("Error while generating qr: ",err)
		return "", "", err
	}

	qrCodeBase64 := base64.StdEncoding.EncodeToString(qrCode)
	
	return  totpSecret,qrCodeBase64,nil
}

func (userData *UserServiceModel)ValidateUser (Login *models.Login) (err error){
	log.Println("Data entered by user to be validated",Login)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result:= userData.UserCollection.FindOne(ctx,bson.M{
        "UserName":Login.UserName,
        "Password":Login.Password,
    })

	if result.Err() != nil {
		log.Println("Error validating user: ", err)
		error := errors.New("wrong User name and password")
		return error
	} else {
		log.Println("User found successfully")
	}

	var validatedUser models.User

	if err := result.Decode(&validatedUser); err != nil {
		log.Println("Error decoding",err)
	}


	log.Println("Successfully Validated")
	log.Println("The user is ",validatedUser)
	return nil
}

func(userData *UserServiceModel)InsertIP(ip *models.IPAddress) (err error){
	log.Println("IP address to be added: ",ip)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	check := userData.IPCollection.FindOne(ctx,ip)

	if check.Err() == nil {
		log.Println("The IP you are trying to insert already exists")
		error := errors.New("ip already exists")
		return error
	}

	result, err := userData.IPCollection.InsertOne(ctx, ip)

	if err != nil {
		log.Println("Error inserting to MongoDB: ", err)
		return err
	} else {
		log.Println("IP Inserted to MongoDb successfully")
	}

	log.Println("Successfully inserted",result)

	return nil
}

func (userData *UserServiceModel) ValidateIP(ip *models.IPAddress) (err error){
	log.Println("IP address of the user: ",ip)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result:= userData.IPCollection.FindOne(ctx,ip)

	if result.Err() != nil {
		log.Println("Error validating IP: ", err)
		error := errors.New("wrong IP Address")
		return error
	} else {
		log.Println("IP Address validated successfully")
	}

	return nil
}

func (userData *UserServiceModel) ValidateTotp(user *models.Login) (company string,role string,plantType string,err error){
	log.Println("Totp entered by the user:",user.UserName,",is : ",user.Totp)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result:= userData.UserCollection.FindOne(ctx,bson.M{"UserName": user.UserName})

	if result.Err() != nil {
		log.Println("Error validating TOTP: ", err)
		error := errors.New("wrong OTP entered")
		return "", "", "",error
	}

	var secretuser models.User

	if err := result.Decode(&secretuser); err != nil {
		log.Println("Error finding Totp code")
	}

	secret := user.Totp
	valid := totp.Validate(secret,secretuser.Totp)

	if !valid {
		log.Println("Wrong TOTP entered")
		err := errors.New("wrong OTP entered")
		return "", "", "",err
	}

	log.Println("Company of the user ",secretuser.Company)

	return secretuser.Company,secretuser.Role,secretuser.PowerplantType,nil
}

func (userData *UserServiceModel) DisplayUser() (allusers *[]models.ShowUser,err error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result,err:=userData.UserCollection.Find(ctx,bson.D{})

	if err != nil {
		log.Println("Error finding data in MongoDB: ", err)
		return nil,err
	} else {
		log.Println("Found data in MongoDb successfully")
	}
	
	var users []models.ShowUser

	result.All(ctx,&users)

	log.Println("All users returned successfully")

	return &users,nil
}