package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/el-Mike/gochat/models"
	"github.com/el-Mike/gochat/persist"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// AuthManager - manages auth related operations.
type AuthManager struct {
	redis *redis.Client
}

var ctx = context.Background()

// NewAuthManager - AuthManager constructor func.
func NewAuthManager() *AuthManager {
	return &AuthManager{
		redis: persist.RedisClient,
	}
}

// Login - authenticates a user.
func (am *AuthManager) Login(user *models.UserModel, apiSecret string) (string, error) {
	authUUID := uuid.New().String()
	userID := user.ID.String()
	email := user.Email
	role := user.Role
	expiresAt := time.Now().Add(time.Minute * 15).Unix()

	token, err := CreateToken(authUUID, userID, email, role, apiSecret, expiresAt)

	if err != nil {
		return "", err
	}

	// Saving authorization allows us to double check the token - when user logs out,
	// token will be removed, and no one will be able to use it anymore, even if it's not
	// expired.
	err = am.redis.Set(ctx, authUUID, userID, 0).Err()

	if err != nil {
		return "", err
	}

	return token, nil
}

// Logout - logs user out by removing it's authorization entry from Redis store.
func (am *AuthManager) Logout(authUUID string) error {
	return am.redis.Del(ctx, authUUID).Err()
}

// VerifyToken - verifies and parses JWT token.
func (am *AuthManager) VerifyToken(request *http.Request, apiSecret string) (*jwt.Token, error) {
	tokenString := am.ExtractToken(request)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(apiSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

// ExtractToken - extracts bearer token from request's headers.
func (am *AuthManager) ExtractToken(request *http.Request) string {
	token := request.Header.Get("Authorization")

	parts := strings.Split(token, " ")

	if len(parts) == 2 {
		return parts[1]
	}

	return ""
}

// HashAndSalt - returned bcrypt hashed password.
func (am *AuthManager) HashAndSalt(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// ComparePasswords - checks if password matches it's hased version.
func (am *AuthManager) ComparePasswords(hashedPassword string, plainPassword []byte) error {
	byteHash := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)

	if err != nil {
		return err
	}

	return nil
}
