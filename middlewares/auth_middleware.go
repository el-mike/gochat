package middlewares

import (
	"errors"
	"net/http"
	"os"

	"github.com/google/uuid"

	"github.com/dgrijalva/jwt-go"
	"github.com/el-Mike/gochat/auth"
	"github.com/el-Mike/gochat/common/api"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware - middleware for authenticating and authorizing the user.
func AuthMiddleware() gin.HandlerFunc {
	authManager := auth.NewAuthManager()

	return func(ctx *gin.Context) {
		request := ctx.Request

		apiSecret := os.Getenv("API_SECRET")

		token, err := authManager.VerifyToken(request, apiSecret)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, api.FromError(err))
			ctx.Abort()

			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, api.FromError(errors.New("Please log in again")))
			ctx.Abort()

			return
		}

		authUUID, authUUIDErr := uuid.Parse(claims["authUUID"].(string))
		userID, userIDErr := uuid.Parse(claims["userID"].(string))

		if authUUIDErr != nil || userIDErr != nil {
			ctx.JSON(http.StatusUnauthorized, api.FromError(errors.New("Token malformed")))
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