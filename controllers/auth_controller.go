package controllers

import (
	"github.com/el-Mike/gochat/auth"
	"github.com/el-Mike/gochat/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		authService: services.NewAuthService(),
	}
}

func (ac *AuthController) SignUp(c *gin.Context) {
	var credentials auth.Credentials

	c.BindJSON(&credentials)

	ac.authService.SignUp(credentials)
}
