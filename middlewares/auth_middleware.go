package middlewares

import (
	"os"

	"github.com/google/uuid"

	"github.com/dgrijalva/jwt-go"
	"github.com/el-Mike/gochat/auth"
	"github.com/el-Mike/gochat/common/api"
	"github.com/el-Mike/gochat/persist"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware - middleware for authenticating and authorizing the user.
func AuthMiddleware() gin.HandlerFunc {
	authManager := auth.NewAuthManager()
	redis := persist.RedisClient

	return func(ctx *gin.Context) {
		request := ctx.Request
		apiSecret := os.Getenv("API_SECRET")
		token, err := authManager.VerifyToken(request, apiSecret)

		if err != nil {
			ctx.JSON(api.ResponseFromError(api.NewAuthorizationError(err)))
			ctx.Abort()

			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			ctx.JSON(api.ResponseFromError(api.NewTokenExpiredError()))
			ctx.Abort()

			return
		}

		authUUID, authUUIDErr := uuid.Parse(claims["authUUID"].(string))
		userID, userIDErr := uuid.Parse(claims["userID"].(string))

		if authUUIDErr != nil || userIDErr != nil {
			ctx.JSON(api.ResponseFromError(api.NewTokenMalforedError()))
			ctx.Abort()

			return
		}

		// If there is no authorization entry is Redis store, it means that
		// user logged out from the application - therefore token expired,
		// even if it's still valid time-wise.
		if err := redis.Get(ctx, authUUID.String()).Err(); err != nil {
			ctx.JSON(api.ResponseFromError(api.NewTokenExpiredError()))
			ctx.Abort()

			return
		}

		email := claims["email"].(string)

		currentUser := &api.ContextUser{
			ID:       userID,
			AuthUUID: authUUID,
			Email:    email,
		}

		ctx.Set(api.ContextUserKey, currentUser)
		ctx.Next()
	}
}
