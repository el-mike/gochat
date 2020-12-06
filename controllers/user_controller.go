package controllers

import (
	"fmt"

	"github.com/el-Mike/gochat/models"
	"github.com/el-Mike/gochat/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) GetUser(c *gin.Context) {}

func (uc *UserController) GetUsers(c *gin.Context) {
	var users []models.User

	if err := uc.userService.GetUsers(&users); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, users)
	}
}

func (uc *UserController) SaveUser(c *gin.Context) {
	var user models.User

	c.BindJSON(&user)

	if err := uc.userService.SaveUser(&user); err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}
}
