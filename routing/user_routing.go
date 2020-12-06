package routing

import (
	"github.com/el-Mike/gochat/controllers"
	"github.com/gin-gonic/gin"
)

func DefineUserRoutes(router *gin.Engine) {
	userController := controllers.NewUserController()

	router.GET("/users", userController.GetUsers)
	router.POST("/users", userController.SaveUser)
}
