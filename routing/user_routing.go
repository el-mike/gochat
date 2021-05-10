package routing

import (
	"github.com/el-Mike/gochat/controllers"
	"github.com/el-Mike/gochat/core/control"
	"github.com/el-Mike/gochat/core/control/rbac"
	"github.com/el-Mike/gochat/models"
	"github.com/gin-gonic/gin"
)

// DefineUserRoutes - registers user routes.
func DefineUserRoutes(router *gin.RouterGroup) {
	handlerCreator := control.NewHandlerCreator()
	userController := controllers.NewUserController()

	// Authenticated routes
	router.GET("/me", handlerCreator.CreateAuthenticated(
		userController.GetMe,
		[]*control.AccessRule{},
	))

	router.GET("/", handlerCreator.CreateAuthenticated(
		userController.GetUsers,
		[]*control.AccessRule{
			{
				ResourceID: models.USER_RESOURCE,
				Action:     rbac.ReadAny,
			},
		},
	))
	router.POST("/", handlerCreator.CreateAuthenticated(
		userController.SaveUser,
		[]*control.AccessRule{
			{
				ResourceID: models.USER_RESOURCE,
				Action:     rbac.Create,
			},
		},
	))
	router.DELETE("/:id", handlerCreator.CreateAuthenticated(
		userController.DeleteUser,
		[]*control.AccessRule{
			{
				ResourceID: models.USER_RESOURCE,
				Action:     rbac.DeleteAny,
			},
		},
	))
}
