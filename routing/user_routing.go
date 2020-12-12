package routing

import (
	"github.com/el-Mike/gochat/controllers"
	"github.com/gin-gonic/gin"
)

func DefineUserRoutes(router *gin.RouterGroup) {
	userController := controllers.NewUserController()

	router.GET("/", userController.GetUsers)
	router.POST("/", userController.SaveUser)
}
