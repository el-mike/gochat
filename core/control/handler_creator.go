package control

import (
	"log"
	"os"

	"github.com/el-Mike/gochat/core/api"
	"github.com/el-Mike/gochat/core/control/rbac"
	"github.com/gin-gonic/gin"
)

// BasicControllerFn - controller function for unauthenticated routes.
type BasicControllerFn func(
	ctx *gin.Context,
) (interface{}, *api.APIError)

// AuthenticatedControllerFn - controller function for authenticated routes.
type AuthenticatedControllerFn func(
	ctx *gin.Context,
	user *ContextUser,
) (interface{}, *api.APIError)

// HandlerCreator - takes desired controller function and produces
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
		accessManager: rbac.AM,
	}
}

// CreateUnauthenticated - creates unauthenticated route.
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

// CreateAuthenticated - creates authenticated route.
func (hc *HandlerCreator) CreateAuthenticated(
	controllerFn AuthenticatedControllerFn,
	accessRules []*AccessRule,
) gin.HandlerFunc {
	apiSecret := os.Getenv("API_SECRET")

	return func(ctx *gin.Context) {
		contextUser, err := hc.authGuard.CheckAuth(ctx.Request, apiSecret)

		if err != nil {
			ctx.JSON(api.ResponseFromError(err))
			return
		}

		for _, rule := range accessRules {

			if rule.Action == "" || rule.ResourceID == "" {
				log.Print("Malformed AccessRule - omitting...")
				continue
			}

			canAccess, err := hc.accessManager.IsGranted(contextUser.Role, rule.ResourceID, rule.Action)

			if err != nil {
				ctx.JSON(api.ResponseFromError(api.NewInternalError(err)))
				return
			}

			if !canAccess {
				ctx.JSON(api.ResponseFromError(api.NewAccessDeniedError(rule.ResourceID, string(rule.Action))))
				return
			}
		}

		result, err := controllerFn(ctx, contextUser)

		if err != nil {
			ctx.JSON(api.ResponseFromError(err))
			return
		}

		ctx.JSON(api.GetSuccessResponse(result))
	}
}
