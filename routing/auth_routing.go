package routing

import (
	"github.com/el-Mike/gochat/controllers"
	"github.com/gin-gonic/gin"
)

func DefineAuthRoutes(router *gin.Engine) {
	authController := controllers.NewAuthController()

	router.POST("/signup", authController.SignUp)
}
