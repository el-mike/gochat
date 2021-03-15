package control

import (
	"os"

	"github.com/el-Mike/gochat/common/api"
	"github.com/gin-gonic/gin"
)

// Controller function for unauthenticated routes.
type BasicControllerFn func(
	ctx *gin.Context,
) (interface{}, *api.APIError)

// Controller function for authenticated routes.
type AuthenticatedControllerFn func(
	ctx *gin.Context,
	user *ContextUser,
) (interface{}, *api.APIError)

// HandlerCreator takes desired controller function and produces
// gin's HandlerFunc. It also takes care of setting response body based on
// controller's return values.
type HandlerCreator struct {
	authGuard *AuthGuard
}

// Returns HandlerCreator instance.
func NewHandlerCreator() *HandlerCreator {
	return &HandlerCreator{
		authGuard: NewAuthGuard(),
	}
}

// Creates unauthenticated route.
func (hc *HandlerCreator) CreateUnauthenticated(controllerFn BasicControllerFn) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		result, err := controllerFn(ctx)

		if err != nil {
			ctx.JSON(api.ResponseFromError(err))
			return
		}

		ctx.JSON(api.GetSuccessResponse(result))
	}
}

// Creates authenticated route.
func (hc *HandlerCreator) CreateAuthenticated(controllerFn AuthenticatedControllerFn) gin.HandlerFunc {
	apiSecret := os.Getenv("API_SECRET")

	return func(ctx *gin.Context) {
		contextUser, err := hc.authGuard.CheckAuth(ctx.Request, apiSecret)

		if err != nil {
			ctx.JSON(api.ResponseFromError(err))
			return
		}

		result, err := controllerFn(ctx, contextUser)

		if err != nil {
			ctx.JSON(api.ResponseFromError(err))
			return
		}

		ctx.JSON(api.GetSuccessResponse(result))
	}
}
