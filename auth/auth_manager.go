package auth

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/el-Mike/gochat/common/interfaces"
	"github.com/el-Mike/gochat/models"
	"github.com/el-Mike/gochat/persist"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type jwtProvider interface {
	CreateToken(authUUID, userID, email, role, apiSecret string, time int64) (string, error)
	ParseToken(tokenString string, apiSecret string) (*jwt.Token, error)
}

type cryptoProvider interface {
	GenerateFromPassword(password []byte, cost int) ([]byte, error)
	CompareHashAndPassword(hashedPassword, password []byte) error
}

type bcryptDelegate struct{}

func (bw *bcryptDelegate) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, cost)
}

func (bw *bcryptDelegate) CompareHashAndPassword(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

// AuthManager - manages auth related operations.
type AuthManager struct {
	redis  interfaces.RedisCache
	jwt    jwtProvider
	crypto cryptoProvider
	ctx    context.Context
}

// NewAuthManager - AuthManager constructor func.
func NewAuthManager() *AuthManager {
	return &AuthManager{
		redis:  *persist.RedisClient,
		jwt:    NewJWTManager(),
		crypto: &bcryptDelegate{},
		ctx:    context.Background(),
	}
}

// Login - authenticates a user.
func (am *AuthManager) Login(user *models.UserModel, apiSecret string) (string, error) {
	authUUID := uuid.New().String()
	userID := user.ID.String()
	email := user.Email
	role := user.Role
	expiresAt := time.Now().Add(time.Minute * 15).Unix()

	token, err := am.jwt.CreateToken(authUUID, userID, email, role, apiSecret, expiresAt)

	if err != nil {
		return "", err
	}

	// Saving authorization allows us to double check the token - when user logs out,
	// token will be removed, and no one will be able to use it anymore, even if it's not
	// expired.
	err = am.redis.Set(am.ctx, authUUID, userID, 0).Err()

	if err != nil {
		return "", err
	}

	return token, nil
}

// Logout - logs user out by removing it's authorization entry from Redis store.
func (am *AuthManager) Logout(authUUID string) error {
	return am.redis.Del(am.ctx, authUUID).Err()
}

// VerifyToken - verifies and parses JWT token.
func (am *AuthManager) VerifyToken(request *http.Request, apiSecret string) (*jwt.Token, error) {
	tokenString := am.extractToken(request)

	token, err := am.jwt.ParseToken(tokenString, apiSecret)

	if err != nil {
		return nil, err
	}

	return token, nil
}

// extractToken - extracts bearer token from request's headers.
func (am *AuthManager) extractToken(request *http.Request) string {
	token := request.Header.Get("Authorization")

	parts := strings.Split(token, " ")

	if len(parts) == 2 {
		return parts[1]
	}

	return ""
}

// HashAndSalt - returned bcrypt hashed password.
func (am *AuthManager) HashAndSalt(password []byte) (string, error) {
	hash, err := am.crypto.GenerateFromPassword(password, bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// ComparePasswords - checks if password matches it's hased version.
func (am *AuthManager) ComparePasswords(hashedPassword string, plainPassword []byte) error {
	byteHash := []byte(hashedPassword)

	err := am.crypto.CompareHashAndPassword(byteHash, plainPassword)

	if err != nil {
		return err
	}

	return nil
}
