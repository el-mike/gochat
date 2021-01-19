package routing

import (
	"github.com/el-Mike/gochat/controllers"
	"github.com/el-Mike/gochat/middlewares"
	"github.com/gin-gonic/gin"
)

// DefineUserRoutes - registers user routes.
func DefineUserRoutes(router *gin.RouterGroup) {
	userController := controllers.NewUserController()

	router.Use(middlewares.AuthMiddleware())

	router.GET("/me", userController.GetMe)

	router.GET("/", userController.GetUsers)
	router.POST("/", userController.SaveUser)
}
