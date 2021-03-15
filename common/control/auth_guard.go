package control

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/el-Mike/gochat/auth"
	"github.com/el-Mike/gochat/common/api"
	"github.com/el-Mike/gochat/persist"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

var ctx = context.Background()

// AuthGuard checks if given request can be properly authenticated, by
// veryfying the token.
type AuthGuard struct {
	authManager *auth.AuthManager
	redis       *redis.Client
}

// Returns new AuthGuard instance.
func NewAuthGuard() *AuthGuard {
	return &AuthGuard{
		authManager: auth.NewAuthManager(),
		redis:       persist.RedisClient,
	}
}

// Checks if given request contains valid token, and returns ContextUser if so.
// Otherwise, APIError will be returned.
func (ag *AuthGuard) CheckAuth(request *http.Request, apiSecret string) (*ContextUser, *api.APIError) {
	token, err := ag.authManager.VerifyToken(request, apiSecret)

	if err != nil {
		return nil, api.NewAuthorizationError(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, api.NewTokenExpiredError()
	}

	authUUID, authUUIDErr := uuid.Parse(claims["authUUID"].(string))
	userID, userIDErr := uuid.Parse(claims["userID"].(string))

	if authUUIDErr != nil || userIDErr != nil {
		return nil, api.NewTokenMalforedError()
	}

	// If there is no authentication entry is Redis store, it means that
	// user logged out from the application - therefore token expired,
	// even if it's still valid time-wise.
	if err := ag.redis.Get(ctx, authUUID.String()).Err(); err != nil {
		return nil, api.NewTokenExpiredError()
	}

	email := claims["email"].(string)

	currentUser := &ContextUser{
		ID:       userID,
		AuthUUID: authUUID,
		Email:    email,
	}

	return currentUser, nil
}
