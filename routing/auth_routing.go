package routing

import (
	"github.com/el-Mike/gochat/controllers"
	"github.com/el-Mike/gochat/core/control"
	"github.com/gin-gonic/gin"
)

// DefineAuthRoutes - registers auth routes.
func DefineAuthRoutes(router *gin.RouterGroup) {
	handlerCreator := control.NewHandlerCreator()
	authController := controllers.NewAuthController()

	// Unauthenticated routes
	router.POST("/signup", handlerCreator.CreateUnauthenticated(authController.SignUp))
	router.POST("/login", handlerCreator.CreateUnauthenticated(authController.Login))

	// Authenticated routes
	router.POST("/logout", handlerCreator.CreateAuthenticated(
		authController.Logout,
		[]*control.AccessRule{},
	))
}
