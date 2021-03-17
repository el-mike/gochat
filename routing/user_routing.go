package routing

import (
	"github.com/el-Mike/gochat/common/control"
	"github.com/el-Mike/gochat/common/control/rbac"
	"github.com/el-Mike/gochat/controllers"
	"github.com/gin-gonic/gin"
)

// DefineUserRoutes - registers user routes.
func DefineUserRoutes(router *gin.RouterGroup) {
	handlerCreator := control.NewHandlerCreator()
	userController := controllers.NewUserController()

	// Authenticated routes
	router.GET("/me", handlerCreator.CreateAuthenticated(
		userController.GetMe,
		[]*rbac.Permission{},
	))

	router.GET("/", handlerCreator.CreateAuthenticated(
		userController.GetUsers,
		[]*rbac.Permission{},
	))
	router.POST("/", handlerCreator.CreateAuthenticated(
		userController.SaveUser,
		[]*rbac.Permission{},
	))
}
