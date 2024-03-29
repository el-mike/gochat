package control

import (
	"log"
	"os"

	"github.com/el-Mike/gochat/core/api"
	"github.com/el-Mike/restrict"
	"github.com/el-Mike/restrict/adapters"
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
	accessManager *restrict.AccessManager
}

// NewHandlerCreator - returns HandlerCreator instance.
func NewHandlerCreator() (*HandlerCreator, error) {
	policyManager, err := restrict.NewPolicyManager(adapters.NewInMemoryAdapter(Policy), true)
	if err != nil {
		return nil, err
	}

	return &HandlerCreator{
		authGuard:     NewAuthGuard(),
		accessManager: restrict.NewAccessManager(policyManager),
	}, nil
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

			var resource restrict.Resource

			if rule.ResourceProvider != nil {
				resource = rule.ResourceProvider(ctx, contextUser)
			} else {
				resource = restrict.UseResource(rule.ResourceID)
			}

			err := hc.accessManager.Authorize(&restrict.AccessRequest{
				Subject:  contextUser,
				Resource: resource,
				Actions:  []string{rule.Action},
			})

			if _, ok := err.(*restrict.AccessDeniedError); err != nil && ok {
				ctx.JSON(api.ResponseFromError(api.NewAccessDeniedError(rule.ResourceID, string(rule.Action))))
				return
			} else {
				ctx.JSON(api.ResponseFromError(api.NewInternalError(err)))
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
