package controllers

import (
	"PowerPlantManagementApplication/interfaces"
	"PowerPlantManagementApplication/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{
	UserService interfaces.Iuser
}

func InitUserController (userService interfaces.Iuser) (UserController){
	return UserController{
		UserService: userService,
	}
}

func (controller *UserController)CreateUser (c *gin.Context){
	var user *models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"Error":"Invalid JSON format",
		})
		return
	}

	
	if qr,err := controller.UserService.CreateUser(user); err != nil {
		c.JSON(http.StatusConflict,gin.H{
			"Error":err.Error(),
		})
		return
	}else{
		c.JSON(http.StatusOK,gin.H{
			"QR":qr,
			"Status":"Success user created successfully",
			"Inserted":user,
		})
	}
}

func (controller *UserController)InsertIP(c *gin.Context){
	var ip *models.IPAddress

	if err := c.ShouldBindJSON(&ip); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"Error":"Invalid JSON format",
		})
		return
	}

	if err := controller.UserService.InsertIP(ip); err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"Error":err,
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"Status":"Success IP inserted successfully",
		"Inserted":ip,
	})

}

func (controller *UserController)ValidateUser(c *gin.Context){
	type logindata struct{
    models.Login `json:"user"`
	models.IPAddress `json:"ip"`
	}
	
	var data logindata

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"Error":"Invalid JSON format",
		})
		return
	}

	if err := controller.UserService.ValidateIP(&data.IPAddress); err != nil {
		c.JSON(http.StatusUnauthorized,gin.H{
			"Error":"Unauthorised IP Address",
		})
		return
	}

	if err := controller.UserService.ValidateUser(&data.Login); err != nil {
		c.JSON(http.StatusUnauthorized,gin.H{
			"Error":"Wrong Username and Password",
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"Status":"Success",
		"Message":"Successfully Validated",
	})
}

func (controller *UserController)ValidateTotp(c *gin.Context){
	var Totpcode *models.Login
	if err := c.ShouldBind(&Totpcode); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"Error":"Invalid JSON format",
		})
		return
	}

	if err := controller.UserService.ValidateTotp(Totpcode);err != nil{
		c.JSON(http.StatusUnauthorized,gin.H{
			"Error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"Status":"Success",
		"Message":"Successfully Validated",
	})
}