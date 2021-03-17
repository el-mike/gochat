package control

import (
	"os"

	"github.com/el-Mike/gochat/common/api"
	"github.com/el-Mike/gochat/common/control/rbac"
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
	authGuard     *AuthGuard
	accessManager *rbac.AccessManager
}

// Returns HandlerCreator instance.
func NewHandlerCreator() *HandlerCreator {
	return &HandlerCreator{
		authGuard:     NewAuthGuard(),
		accessManager: rbac.NewAccessManager(),
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
func (hc *HandlerCreator) CreateAuthenticated(
	controllerFn AuthenticatedControllerFn,
	requiredPermissions []*rbac.Permission,
) gin.HandlerFunc {
	apiSecret := os.Getenv("API_SECRET")

	return func(ctx *gin.Context) {
		contextUser, err := hc.authGuard.CheckAuth(ctx.Request, apiSecret)

		if err != nil {
			ctx.JSON(api.ResponseFromError(err))
			return
		}

		role := contextUser.Role()

		if canAccess := hc.accessManager.IsGranted(role, requiredPermissions...); !canAccess {
			ctx.JSON(api.ResponseFromError(api.NewAccessDeniedError()))
		}

		result, err := controllerFn(ctx, contextUser)

		if err != nil {
			ctx.JSON(api.ResponseFromError(err))
			return
		}

		ctx.JSON(api.GetSuccessResponse(result))
	}
}
